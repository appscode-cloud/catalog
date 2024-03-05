# catalog

## Project Generator Commands

```bash
> kubebuilder init --domain bytebuilders.dev --skip-go-version-check

> kubebuilder create api --resource=true --controller=false --force --group catalog --version v1alpha1 --kind DruidBinding
> kubebuilder create api --resource=true --controller=false --force --group catalog --version v1alpha1 --kind ElasticsearchBinding
> kubebuilder create api --resource=true --controller=false --force --group catalog --version v1alpha1 --kind FerretDBBinding
> kubebuilder create api --resource=true --controller=false --force --group catalog --version v1alpha1 --kind KafkaBinding
> kubebuilder create api --resource=true --controller=false --force --group catalog --version v1alpha1 --kind MariaDBBinding
> kubebuilder create api --resource=true --controller=false --force --group catalog --version v1alpha1 --kind MemcachedBinding
> kubebuilder create api --resource=true --controller=false --force --group catalog --version v1alpha1 --kind MemcachedBinding
> kubebuilder create api --resource=true --controller=false --force --group catalog --version v1alpha1 --kind MicrosoftSQLServerBinding
> kubebuilder create api --resource=true --controller=false --force --group catalog --version v1alpha1 --kind MongoDBBinding
> kubebuilder create api --resource=true --controller=false --force --group catalog --version v1alpha1 --kind MySQLBinding
> kubebuilder create api --resource=true --controller=false --force --group catalog --version v1alpha1 --kind PerconaXtraDBBinding
> kubebuilder create api --resource=true --controller=false --force --group catalog --version v1alpha1 --kind PgBouncerBinding
> kubebuilder create api --resource=true --controller=false --force --group catalog --version v1alpha1 --kind PgpoolBinding
> kubebuilder create api --resource=true --controller=false --force --group catalog --version v1alpha1 --kind PostgresBinding
> kubebuilder create api --resource=true --controller=false --force --group catalog --version v1alpha1 --kind ProxySQLBinding
> kubebuilder create api --resource=true --controller=false --force --group catalog --version v1alpha1 --kind ProxySQLBinding
> kubebuilder create api --resource=true --controller=false --force --group catalog --version v1alpha1 --kind RabbitMQBinding
> kubebuilder create api --resource=true --controller=false --force --group catalog --version v1alpha1 --kind RedisBinding
> kubebuilder create api --resource=true --controller=false --force --group catalog --version v1alpha1 --kind SinglestoreBinding
> kubebuilder create api --resource=true --controller=false --force --group catalog --version v1alpha1 --kind SolrBinding
> kubebuilder create api --resource=true --controller=false --force --group catalog --version v1alpha1 --kind ZooKeeperBinding
```

## Getting Started
You’ll need a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) to get a local cluster for testing, or run against a remote cluster.
**Note:** Your controller will automatically use the current context in your kubeconfig file (i.e. whatever cluster `kubectl cluster-info` shows).

### Install CRDs
To install the CRDs into the cluster:

```sh
make install
```

### Uninstall CRDs
To delete the CRDs from the cluster:

```sh
make uninstall
```

### Modifying the API definitions
If you are editing the API definitions, generate the manifests such as CRs or CRDs using:

```sh
make manifests
```

**NOTE:** Run `make --help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

