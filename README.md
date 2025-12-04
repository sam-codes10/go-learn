## Auth-Service

This is auth service. It is used to authenticate and authorize. It is useful for preventing unauthorized access to other services or api-endpoints.

There are following endpoints in a typical auth-service : 
- /signup
- /login
- /reset-password
- /verify-otp

For sending OTPs, the auth-service generates the OTP and stores it in redis with a 10 minutes TTL, then the auth-service calls the mailer-service and the mailer service sends the OTP on the mail or the phone number.