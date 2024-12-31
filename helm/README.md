# school-bot

Add chart

```bash
helm repo add linuxoid69 https://linuxoid69.github.io/helm-charts
helm repo update
helm upgrade --install -n bots school-bot -f school-bot-values.yaml linuxoid69/school-bot --version 0.1.0
kubectl -n bots rollout restart deployment school-bot
```
