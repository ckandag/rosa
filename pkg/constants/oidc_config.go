package constants

const (
	InstallerRoleArnFlag = "installer-role-arn"
	IssuerUrlFlag        = "issuer-url"
	SecretArnFlag        = "secret-arn"

	SecretsManagerService = "secretsmanager"

	MinorVersionForGetSecret  = "4.12"
	InformOperatorRolesOutput = "To create Operator Roles for this OIDC Configuration, " +
		"run the following command and remember to replace <user-defined> with a prefix of your choice:\n" +
		"\trosa create operator-roles --prefix <user-defined> --oidc-config-id %s\n" +
		"If you are going to create a Hosted Control Plane cluster please include '--hosted-cp'"
)
