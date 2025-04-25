9. Security & Authentication

Keeping systems secure requires layered controls—from authenticating API calls to protecting data in flight and at rest. Below, each topic is broken down with simple examples.
9.1 API Authentication
OAuth 2.0

    What it is: An authorization framework where a client (e.g., web or mobile app) obtains an access token to act on behalf of a user.

    Roles:

        Resource Owner: The user.

        Client: Your application.

        Authorization Server: Issues tokens.

        Resource Server: Hosts protected APIs.

    Flow (Authorization Code Grant):

        User logs into Auth Server and consents.

        Client receives an authorization code.

        Client exchanges code (plus secret) for an access token.

        Client calls Resource Server with Authorization: Bearer <token>.

    Example:

        Browser → GET /authorize?response_type=code&client_id=xyz.

        Auth server → redirect back with ?code=ABC.

        Client → POST /token (code, secret) → receives { access_token, refresh_token }.

        Client → GET /orders with Authorization: Bearer access_token.

OpenID Connect (OIDC)

    What it is: A thin identity layer on top of OAuth 2.0 that returns an ID token (a JWT) containing user profile claims (e.g., email, name).

    Use Case: Authenticate users (prove identity) as well as authorize.

JSON Web Tokens (JWT)

    Structure:

        Header (algorithm, type)

        Payload (claims: sub, exp, custom data)

        Signature (HMAC or RSA)

    Why: Self-contained; no need to lookup session in DB for every request.

    Example Payload:

    {
      "sub": "user123",
      "iss": "https://auth.example.com",
      "exp": 1714000000,
      "roles": ["order:create","order:view"]
    }

    Validation:

        Verify signature with shared secret or public key.

        Check exp (expiry) and iss (issuer).

        Inspect claims for authorization.

9.2 Network Security
Mutual TLS (mTLS)

    What it is: Both client and server present certificates and verify each other.

    Use Case: Secure service-to-service calls in a microservices mesh.

    Flow:

        Client presents its cert.

        Server checks client cert against trusted CA.

        Server presents its cert; client verifies it.

        Encrypted channel established.

Zero-Trust & IAM

    Zero-Trust Principle: “Never trust; always verify.” No implicit network trust—every request is authenticated and authorized.

    IAM (Identity & Access Management):

        Centralizes user/service identities, roles, and permissions.

        Examples: AWS IAM, Google Cloud IAM, Azure AD.

    Example Policy (AWS IAM):

    {
      "Version": "2012-10-17",
      "Statement": [{
        "Effect": "Allow",
        "Action": ["dynamodb:Query", "dynamodb:GetItem"],
        "Resource": "arn:aws:dynamodb:us-east-1:123456789012:table/Orders"
      }]
    }

9.3 Rate Limiting & DDoS Protections
Web Application Firewalls (WAF)

    Purpose: Inspect incoming HTTP(s) traffic for malicious patterns (SQL injection, XSS).

    Example: AWS WAF or Cloudflare WAF blocking requests with suspicious payloads.

DDoS Protection

    AWS Shield (Standard & Advanced)

        Automatically protects AWS resources (ELB, CloudFront) against volumetric attacks.

    Cloudflare

        Global Anycast network absorbs and filters malicious traffic before it hits your origin.

How It Works Together

    WAF filters out bad requests (layer 7).

    Rate limits throttle excessive requests per IP or API key.

    DDoS network (AWS Shield/Cloudflare) scrubs high-volume attacks (layer 3/4).

9.4 Data Security
Encryption At-Rest

    What it is: Data stored on disk or in databases is encrypted.

    Key Management:

        KMS (Key Management Service): AWS KMS, Google KMS—manages creation, rotation, and permissions for encryption keys.

        HSM (Hardware Security Module): Dedicated keystores with tamper-proof hardware.

    Example:

        Amazon S3 bucket with server-side encryption using an AWS-managed KMS key.

        RDS database encrypted with a KMS key.

Encryption In-Transit

    What it is: TLS (HTTPS) ensures data is encrypted between client and server (and between services).

    Certificates:

        Provision via public CAs (Let’s Encrypt) or internal PKI.

        Rotate and renew regularly.

Audit Logging

    What it is: Record all security-relevant events (logins, permission changes, data access) for compliance and forensics.

    Examples:

        AWS CloudTrail logs all API calls.

        Database audit logs record every SELECT, INSERT, UPDATE, DELETE with timestamp, user, and query.

Putting It All Together

    Authenticate users and services with OAuth 2.0/OIDC and JWTs.

    Enforce network security using mTLS in service meshes and a zero-trust IAM foundation.

    Protect at the edge with WAF, rate limits, and DDoS defense (Cloudflare, AWS Shield).

    Encrypt data everywhere: at-rest with KMS/HSM, in-transit with TLS.

    Audit all security events to detect anomalies and meet compliance needs.

By layering these controls, you build a defense-in-depth posture that protects your APIs, network, and data against evolving threats.

