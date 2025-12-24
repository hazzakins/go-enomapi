AddDnsSec
=========

This command will set DNS Sec directly at the registry tied to a specific domain.

Usage
-----

This command creates a DS record in the responsible registry's DNS zone as part of a DNSSEC chain of trust.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and API Token must be valid.
- The domain name must belong to this account.

Input Parameters
----------------

> ### You can optionally also set, or only set the "MaxSigLife" parameter. Not all registries will support DNS Sec, and of those that do, not all of them will support setting of "MaxSigLife".
>
>
>
> If the "SetMaxLifeOnly" parameter is passed in, the API command will only attempt to set that value and it will ignore anything else passed in. If you pass that parameter in as "True", then the "MaxSigLife" input parameter is required. Otherwise, you can optionally pass "MaxSigLife" in along with the other DNS Key data for it to be attempted to be set at the registry.

```
https://resellertest.enom.com/interface.asp?command=AddDnsSec&uid=YourAccountID&pw=YourApiToken&SLD={Required}&TLD={Required}&SetMaxLifeOnly={Optional}&MaxSigLife={Optional}&Alg={Required}&Digest{Required}&DigestType={Required}&KeyTag={Required}&responsetype={Optional}
```
| Input Parameter | Type | Status                 | Description |
| --------------- | ------ | -------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Command     | string | Required                | AddDnsSec |
| UID | string | Required | Your Account ID                                                                                                                             |
| PW       | string | Required                | Your API Token |
| SLD | string | Required | Second-level domain name \(e.g. "enom" in "enom.com"\)                                                                                                        |
| TLD       | string | Required                | Top-level domain name \(extension\) |
| SetMaxLifeOnly | bool  | Optional | Permitted values are: - True - False                                                                                                                   |
| MaxSigLife   | int | Required if "SetMaxLifeOnly" is used. | |
| Alg | int  | Required | Algorithm used in generating the Digest value. Permitted values are: - "3" - DSA/SHA-1 - "5" - RSA/SHA-1 - "7" - RSASHA1-NSEC3-SHA1 - "8" - RSA/SHA-256 - "10" - RSA/SHA-512 - "12" - GOST R 34.10-2001 - "13" - ECDSA/SHA-256 - "14" - ECDSA/SHA-384 |
| Digest     | string | Required                | |
| DigestType | int  | Required | Permitted values are: - "1" - SHA-1 - "2" - SHA-256                                                                                                         |
| KeyTag     | int | Required                | |
| ResponseType | string | .Optional | The format that the system respond in. Permitted values are - HTML - XML - TEXT The default response format is "TEXT".                                                                         |

Returned Parameters and Values
------------------------------

Check the return parameter "ErrCount". If it is greater than 0, the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type | Description                                                      |
| ---------------- | ------- | ---------------------------------------------------------------------------------------------------------------------- |
| Command     | string | Name of command executed                                                |
| ErrCount     | int | The number of errors if any occurred. If the value is greater then 0, check the Err\(1 to ErrCount\) values.     |
| Err\(X\)     | string | Error messages explaining the failure. These can be presented "as-is" back to the client.              |
| ResponseCode   | int | Numeric value to indicate success or error of the executed command.                          |
| ResponseMessage | string | Response messages explaining the status of the command response. These can be presented "as-is" back to the client. |
| done       | boolean | "True" value indicates this entire response has reached you successfully.                       |

Related Commands
----------------

[DeleteDnsSec](../docs/deletednssec.md)

[GetDnsSec](../docs/getdnssec.md)