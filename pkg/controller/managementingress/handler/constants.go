package handler

const (
	AppName                   string = "management-ingress"
	ServiceName               string = "icp-management-ingress"
	ConfigName                string = "management-ingress-config"
	IAMTokenService           string = "iam-token-service"
	ServiceAccountName        string = "management-ingress"
	CertName                  string = "management-ingress-cert"
	TLSSecretName             string = "icp-management-ingress-tls-secret"
	RouteName                 string = "cp-console"
	RouteCert                 string = "route-cert"
	RouteSecret               string = "route-tls-secret"
	csPriorityClassName       string = "cs-priority-class"
	ConfigUpdateAnnotationKey string = "management-ingress.operator.k8s.io/config-updated"
	SCCName                   string = "management-ingress-scc"
	BindInfoConfigMap         string = "management-ingress-info"
	PlatformAuthConfigmap     string = "platform-auth-idp"
	PlatformAuthSecret        string = "platform-oidc-credentials"
)