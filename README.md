# cloudrun-sample

## チュートリアル
https://cloud.google.com/run/docs/quickstarts/build-and-deploy/deploy-go-service?hl=ja

## デプロイ
```
gcloud run deploy --region asia-northeast1
```

## APIの公開設定
### 非公開

<!-- ```
gcloud iam service-accounts create renderer-identity
gcloud run deploy renderer \
--service-account renderer-identity \
--no-allow-unauthenticated
```

または、


```
gcloud run deploy --no-allow-unauthenticated
``` -->

### 公開

```
gcloud run services add-iam-policy-binding helloworld --member="allUsers" --role="roles/run.invoker"
```

または、

```
gcloud run deploy --allow-unauthenticated
```

https://cloud.google.com/run/docs/authenticating/public?hl=ja#gcloud