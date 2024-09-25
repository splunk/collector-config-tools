module github.com/splunk/collector-config-tools/cfgschema

go 1.22

toolchain go1.22.6

require (
	github.com/fatih/structtag v1.2.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awss3exporter v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/fileexporter v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kafkaexporter v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sapmexporter v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/splunkhecexporter v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/ackextension v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/basicauthextension v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/headerssetterextension v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextension v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/httpforwarderextension v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/oauth2clientauthextension v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecsobserver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecstaskobserver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/hostobserver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/pprofextension v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/filestorage v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/cumulativetodeltaprocessor v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/groupbyattrsprocessor v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/logstransformprocessor v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstransformprocessor v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/probabilisticsamplerprocessor v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/redactionprocessor v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/activedirectorydsreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/apachereceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsecscontainermetricsreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azuremonitorreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/carbonreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/cloudfoundryreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/collectdreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/elasticsearchreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/filelogreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/fluentforwardreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudpubsubreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/haproxyreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/httpcheckreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jaegerreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jmxreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/journaldreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkametricsreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkareceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mysqlreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nginxreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/oracledbreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/postgresqlreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/rabbitmqreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/receivercreator v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/redisreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sapmreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/signalfxreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/solacereceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/splunkenterprisereceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/splunkhecreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlqueryreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlserverreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sshcheckreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/syslogreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/tcplogreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/udplogreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/vcenterreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/wavefrontreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowseventlogreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowsperfcountersreceiver v0.108.0
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/zipkinreceiver v0.108.0
	github.com/stretchr/testify v1.9.0
	go.opentelemetry.io/collector/component v0.108.1
	go.opentelemetry.io/collector/config/configtls v1.14.1
	go.opentelemetry.io/collector/exporter v0.108.1
	go.opentelemetry.io/collector/exporter/debugexporter v0.108.1
	go.opentelemetry.io/collector/exporter/loggingexporter v0.108.1
	go.opentelemetry.io/collector/exporter/otlpexporter v0.108.1
	go.opentelemetry.io/collector/exporter/otlphttpexporter v0.108.1
	go.opentelemetry.io/collector/extension v0.108.1
	go.opentelemetry.io/collector/extension/ballastextension v0.108.1
	go.opentelemetry.io/collector/extension/zpagesextension v0.108.1
	go.opentelemetry.io/collector/otelcol v0.108.1
	go.opentelemetry.io/collector/processor v0.108.1
	go.opentelemetry.io/collector/processor/batchprocessor v0.108.1
	go.opentelemetry.io/collector/processor/memorylimiterprocessor v0.108.1
	go.opentelemetry.io/collector/receiver v0.108.1
	go.opentelemetry.io/collector/receiver/otlpreceiver v0.108.1
	golang.org/x/mod v0.21.0
	golang.org/x/text v0.18.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	cloud.google.com/go/auth v0.9.1 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.4 // indirect
	cloud.google.com/go/compute/metadata v0.5.0 // indirect
	cloud.google.com/go/iam v1.1.12 // indirect
	cloud.google.com/go/logging v1.11.0 // indirect
	cloud.google.com/go/longrunning v0.5.11 // indirect
	cloud.google.com/go/pubsub v1.42.0 // indirect
	code.cloudfoundry.org/go-diodes v0.0.0-20211115184647-b584dd5df32c // indirect
	code.cloudfoundry.org/go-loggregator v7.4.0+incompatible // indirect
	code.cloudfoundry.org/rfc5424 v0.0.0-20201103192249-000122071b78 // indirect
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4 // indirect
	github.com/99designs/keyring v1.2.2 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.14.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.7.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.10.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor v0.11.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources v1.2.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/storage/azblob v1.4.0 // indirect
	github.com/Azure/go-amqp v1.1.0 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20230124172434-306776ec8161 // indirect
	github.com/AzureAD/microsoft-authentication-library-for-go v1.2.2 // indirect
	github.com/GehirnInc/crypt v0.0.0-20200316065508-bb7000b8a962 // indirect
	github.com/IBM/sarama v1.43.3 // indirect
	github.com/JohnCGriffin/overflow v0.0.0-20211019200055-46fa312c352c // indirect
	github.com/SAP/go-hdb v1.11.3 // indirect
	github.com/apache/arrow/go/v15 v15.0.0 // indirect
	github.com/apache/thrift v0.20.0 // indirect
	github.com/aws/aws-sdk-go v1.55.5 // indirect
	github.com/aws/aws-sdk-go-v2 v1.30.4 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.6.2 // indirect
	github.com/aws/aws-sdk-go-v2/config v1.27.28 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.17.28 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.16.12 // indirect
	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.16.15 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.16 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.16 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.1 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.3.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.11.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.3.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.11.18 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.17.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3 v1.53.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/servicediscovery v1.31.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.22.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.26.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.30.4 // indirect
	github.com/aws/smithy-go v1.20.4 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bmatcuk/doublestar/v4 v4.6.1 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cloudfoundry-incubator/uaago v0.0.0-20190307164349-8136b7bbe76e // indirect
	github.com/danieljoos/wincred v1.2.1 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/distribution/reference v0.6.0 // indirect
	github.com/docker/docker v27.1.1+incompatible // indirect
	github.com/dvsekhvalnov/jose2go v1.6.0 // indirect
	github.com/eapache/go-resiliency v1.7.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20230731223053-c322873962e3 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/emicklei/go-restful/v3 v3.11.3 // indirect
	github.com/expr-lang/expr v1.16.9 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/fxamacker/cbor/v2 v2.7.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/go-openapi/jsonpointer v0.21.0 // indirect
	github.com/go-openapi/jsonreference v0.21.0 // indirect
	github.com/go-openapi/swag v0.23.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/go-test/deep v1.1.1 // indirect
	github.com/go-viper/mapstructure/v2 v2.1.0 // indirect
	github.com/gobwas/glob v0.2.4-0.20181002190808-e7a84e9525fe // indirect
	github.com/goccy/go-json v0.10.3 // indirect
	github.com/godbus/dbus v0.0.0-20190726142602-4481cbc300e2 // indirect
	github.com/gogo/googleapis v1.4.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.1 // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/golang-sql/sqlexp v0.1.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/flatbuffers v24.3.7+incompatible // indirect
	github.com/google/gnostic-models v0.6.8 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/pprof v0.0.0-20240711041743-f6c9dda6c6da // indirect
	github.com/google/s2a-go v0.1.8 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.2 // indirect
	github.com/googleapis/gax-go/v2 v2.13.0 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.20.0 // indirect
	github.com/gsterjov/go-libsecret v0.0.0-20161001094733-a6f4afe4910c // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/hashicorp/hcl v1.0.1-vault-5 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
	github.com/imdario/mergo v0.3.16 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jaegertracing/jaeger v1.60.0 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.7.6 // indirect
	github.com/jcmturner/goidentity/v6 v6.0.1 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.4 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/jonboulle/clockwork v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/klauspost/cpuid/v2 v2.2.7 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.1 // indirect
	github.com/kr/fs v0.1.0 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/leodido/go-syslog/v4 v4.1.0 // indirect
	github.com/leodido/ragel-machinery v0.0.0-20190525184631-5f46317e436b // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/lightstep/go-expohisto v1.0.0 // indirect
	github.com/lufia/plan9stats v0.0.0-20240226150601-1dcf7310316a // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/microsoft/go-mssqldb v1.7.2 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mongodb-forks/digest v1.1.0 // indirect
	github.com/mostynb/go-grpc-compression v1.2.3 // indirect
	github.com/mtibben/percent v0.2.1 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/nginxinc/nginx-prometheus-exporter v0.11.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer v0.108.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/ecsutil v0.108.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/collectd v0.108.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.108.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.108.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/exp/metrics v0.108.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter v0.108.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/kafka v0.108.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/sharedcomponent v0.108.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk v0.108.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/sqlquery v0.108.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchperresourceattr v0.108.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchpersignal v0.108.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata v0.108.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil v0.108.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/sampling v0.108.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza v0.108.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/azure v0.108.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/jaeger v0.108.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/signalfx v0.108.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/zipkin v0.108.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/winperfcounters v0.108.0 // indirect
	github.com/openzipkin/zipkin-go v0.4.3 // indirect
	github.com/philhofer/fwd v1.1.3-0.20240612014219-fbbf4953d986 // indirect
	github.com/pierrec/lz4/v4 v4.1.21 // indirect
	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c // indirect
	github.com/pkg/sftp v1.13.6 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/power-devops/perfstat v0.0.0-20240221224432-82ca36839d55 // indirect
	github.com/prometheus-community/windows_exporter v0.27.2 // indirect
	github.com/prometheus/client_golang v1.20.2 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.55.0 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/redis/go-redis/v9 v9.6.1 // indirect
	github.com/relvacode/iso8601 v1.4.0 // indirect
	github.com/rs/cors v1.11.0 // indirect
	github.com/shirou/gopsutil/v4 v4.24.7 // indirect
	github.com/shoenig/go-m1cpu v0.1.6 // indirect
	github.com/shoenig/test v1.7.1 // indirect
	github.com/signalfx/com_signalfx_metrics_protobuf v0.0.3 // indirect
	github.com/signalfx/sapm-proto v0.14.0 // indirect
	github.com/sijms/go-ora/v2 v2.8.19 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/snowflakedb/gosnowflake v1.11.0 // indirect
	github.com/spf13/cast v1.7.0 // indirect
	github.com/spf13/cobra v1.8.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/tg123/go-htpasswd v1.2.2 // indirect
	github.com/tinylib/msgp v1.2.0 // indirect
	github.com/tklauser/go-sysconf v0.3.13 // indirect
	github.com/tklauser/numcpus v0.7.0 // indirect
	github.com/valyala/fastjson v1.6.4 // indirect
	github.com/vmware/govmomi v0.42.0 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	github.com/zeebo/xxh3 v1.0.2 // indirect
	go.etcd.io/bbolt v1.3.11 // indirect
	go.mongodb.org/atlas v0.36.0 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/collector v0.108.1 // indirect
	go.opentelemetry.io/collector/client v1.14.1 // indirect
	go.opentelemetry.io/collector/component/componentprofiles v0.108.1 // indirect
	go.opentelemetry.io/collector/component/componentstatus v0.108.1 // indirect
	go.opentelemetry.io/collector/config/configauth v0.108.1 // indirect
	go.opentelemetry.io/collector/config/configcompression v1.14.1 // indirect
	go.opentelemetry.io/collector/config/configgrpc v0.108.1 // indirect
	go.opentelemetry.io/collector/config/confighttp v0.108.1 // indirect
	go.opentelemetry.io/collector/config/confignet v0.108.1 // indirect
	go.opentelemetry.io/collector/config/configopaque v1.14.1 // indirect
	go.opentelemetry.io/collector/config/configretry v1.14.1 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.108.1 // indirect
	go.opentelemetry.io/collector/config/internal v0.108.1 // indirect
	go.opentelemetry.io/collector/confmap v1.14.1 // indirect
	go.opentelemetry.io/collector/connector v0.108.1 // indirect
	go.opentelemetry.io/collector/consumer v0.108.1 // indirect
	go.opentelemetry.io/collector/consumer/consumerprofiles v0.108.1 // indirect
	go.opentelemetry.io/collector/consumer/consumertest v0.108.1 // indirect
	go.opentelemetry.io/collector/extension/auth v0.108.1 // indirect
	go.opentelemetry.io/collector/featuregate v1.14.1 // indirect
	go.opentelemetry.io/collector/filter v0.108.1 // indirect
	go.opentelemetry.io/collector/internal/globalgates v0.108.1 // indirect
	go.opentelemetry.io/collector/pdata v1.14.1 // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.108.1 // indirect
	go.opentelemetry.io/collector/pdata/testdata v0.108.1 // indirect
	go.opentelemetry.io/collector/semconv v0.108.1 // indirect
	go.opentelemetry.io/collector/service v0.108.1 // indirect
	go.opentelemetry.io/contrib/config v0.8.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.53.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.53.0 // indirect
	go.opentelemetry.io/contrib/propagators/b3 v1.28.0 // indirect
	go.opentelemetry.io/contrib/zpages v0.53.0 // indirect
	go.opentelemetry.io/otel v1.28.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp v0.4.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v1.28.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp v1.28.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.28.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.28.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.28.0 // indirect
	go.opentelemetry.io/otel/exporters/prometheus v0.50.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v1.28.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.28.0 // indirect
	go.opentelemetry.io/otel/log v0.4.0 // indirect
	go.opentelemetry.io/otel/metric v1.28.0 // indirect
	go.opentelemetry.io/otel/sdk v1.28.0 // indirect
	go.opentelemetry.io/otel/sdk/log v0.4.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.28.0 // indirect
	go.opentelemetry.io/otel/trace v1.28.0 // indirect
	go.opentelemetry.io/proto/otlp v1.3.1 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/crypto v0.26.0 // indirect
	golang.org/x/exp v0.0.0-20240719175910-8a7402abbf56 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/oauth2 v0.22.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/term v0.23.0 // indirect
	golang.org/x/time v0.6.0 // indirect
	golang.org/x/tools v0.23.0 // indirect
	golang.org/x/xerrors v0.0.0-20231012003039-104605ab7028 // indirect
	gonum.org/v1/gonum v0.15.1 // indirect
	google.golang.org/api v0.194.0 // indirect
	google.golang.org/genproto v0.0.0-20240814211410-ddb44dafa142 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240730163845-b1a4ccb954bf // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240823204242-4ba0660f739c // indirect
	google.golang.org/grpc v1.65.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/api v0.31.0 // indirect
	k8s.io/apimachinery v0.31.0 // indirect
	k8s.io/client-go v0.31.0 // indirect
	k8s.io/klog/v2 v2.130.1 // indirect
	k8s.io/kube-openapi v0.0.0-20240228011516-70dd3763d340 // indirect
	k8s.io/utils v0.0.0-20240711033017-18e509b52bc8 // indirect
	sigs.k8s.io/controller-runtime v0.19.0 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.4.1 // indirect
	sigs.k8s.io/yaml v1.4.0 // indirect
)
