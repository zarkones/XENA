If 'Origin' header is present in the HTTP request, it needs to be compared against whitelisted hosts. If supplied value of the origin header doesn't match any whitelisted host, the request must be rejected. This policy should take effect on all back-end services.

Whitelist should be supplied through the CORS_POLICY_ALLOWED_ORIGINS environment variable, or through other means if that's not possible. CORS_POLICY_ALLOWED_ORIGINS is a comma separated list of trusted hosts. Example:

CORS_POLICY_ALLOWED_ORIGINS=http://127.0.0.1:3000,https://mydomain.com