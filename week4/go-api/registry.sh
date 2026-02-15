
kubectl create secret docker-registry stackit-registry \
  --docker-server=registry.onstackit.cloud \
  --docker-username='mariia.rubina13@gmail.com' \
  --docker-password='iCXcos7diXSrd5htectiTYEPde7prxPs' \
  --docker-email='mariia.rubina13@gmail.com'

iCXcos7diXSrd5htectiTYEPde7prxPs

  kubectl create secret generic stackit-registry  --from-file='/Users/mariia.rubina13/credentials/robot$mariia-api+api-robot.json'