package api

import (

	// "log"
	// "os/exec"
	vb "github.com/OT-CONTAINER-KIT/redis-operator/api/common/v1beta2"
	"github.com/OT-CONTAINER-KIT/redis-operator/api/redisreplication/v1beta2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	// "sigs.k8s.io/controller-runtime"
	"k8s.io/apimachinery/pkg/api/resource"
	// "time"
	corev1 "k8s.io/api/core/v1"
	// "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func setOpts(namespace string, name string) *client.ListOptions {
	var opts client.ListOptions
	opts.Namespace = namespace
	if name != "" {
		var sel metav1.LabelSelector
		sel.MatchLabels = make(map[string]string)
		sel.MatchLabels["name"] = name
		labelSel, err := metav1.LabelSelectorAsSelector(&sel)
		if err != nil {
			panic(err.Error())
		}
		opts.LabelSelector = labelSel
	}
	return &opts
}

func setResource(rr *v1beta2.RedisReplication, name string, namespace string, replnum int32) {
	rr.Name = name
	rr.Namespace = namespace
	size := replnum
	rau := int64(1000)
	fsg := int64(1000)
	var psc corev1.PodSecurityContext

	psc.RunAsUser = &rau
	psc.FSGroup = &fsg
	rr.Spec.Size = &size
	rr.APIVersion = "redis.redis.opstreelabs.in/v1beta2"
	rr.Kind = "RedisReplication"
	rr.Spec.KubernetesConfig.Image = "quay.io/opstree/redis:latest"
	rr.Spec.KubernetesConfig.ImagePullPolicy = "IfNotPresent"
	rr.Spec.PodSecurityContext = &psc
	rr.Spec.Storage = &vb.Storage{}
	rr.Spec.Storage.VolumeClaimTemplate.Spec.AccessModes = []corev1.PersistentVolumeAccessMode{
		corev1.ReadWriteOnce,
	}
	rr.Spec.Storage.VolumeClaimTemplate.Spec.Resources.Requests = make(corev1.ResourceList)
	rr.Spec.Storage.VolumeClaimTemplate.Spec.Resources.Requests[corev1.ResourceStorage] = resource.MustParse("1Gi")
	var exporter vb.RedisExporter
	exporter.Image = "quay.io/opstree/redis-exporter:latest"
	exporter.Enabled = false
	rr.Spec.RedisExporter = &exporter
	var service vb.ServiceConfig
	service.ServiceType = "LoadBalancer"
	annotations := map[string]string{
				"service.beta.kubernetes.io/aws-load-balancer-internal": "0.0.0.0/0",
			}
	service.ServiceAnnotations = annotations
	rr.Spec.KubernetesConfig.Service = &service
}
