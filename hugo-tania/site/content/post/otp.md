---
title: otp
servicename: otp
labels: 
- Readme
---
One time password service

# OTP Service

Generate one time passwords (OTP) for any unique id, email or user. Codes are valid for up to 60 seconds.

## cURL


### Otp Generate
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
Generate an OTP (one time pass) code
```shell
> curl 'https://api.m3o.com/otp/Otp/Generate' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "id": "unique id, email or user to generate an OTP for"
};
# Response
{
  "code": "6 digit one time pass code"
}
```


### Otp Validate
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
Validate the OTP code
```shell
> curl 'https://api.m3o.com/otp/Otp/Validate' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "code": "6 digit one time pass code to validate",
  "id": "unique id, email or user for which the code was generated"
};
# Response
{
  "success": true
}
```


