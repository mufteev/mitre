package stix

// Вспомогательная функция для проверки на revoked и deprecated
func isRevokedOrDeprecated(obj map[string]interface{}) bool {
	if revoked, ok := obj["revoked"].(bool); ok && revoked {
		return true
	}

	if deprecated, ok := obj["x_mitre_deprecated"].(bool); ok && deprecated {
		return true
	}

	return false
}
