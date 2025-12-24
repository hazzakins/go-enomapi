DeleteDnsSec
============

This command will remove previously set DNS Sec on an specified domain, directly at the registry.

Usage
-----

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and API Token must be valid.
- The domain name must belong to this account.

> ### You must pass in all four of the DNS Sec parameters and they must match what is set in order for the record to be removed. If you are unclear as to what the other parameters are, you can run the "GetDnsSec" command to return back all of the proper values.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=DeleteDnsSec&uid=YourAccountID&pw=YourApiToken&SLD={Required}&TLD={Required}&Alg={Required}&Digest{Required}&DigestType={Required}&KeyTag={Required}&responsetype={Optional}
```
| Input Parameter | Type | Status   | Description |
| --------------- | ------ | ---------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| command     | string | Required  | DeleteDnsSec |
| uid | string | Required | Your Account ID                                                                                                                             |
| pw       | string | Required  | Your API Token |
| SLD | string | Required | Second-level domain name \(e.g. "enom" in "enom.com"\)                                                                                                        |
| TLD       | string | Required  | Top-level domain name \(extension\) |
| Alg | string | Required | Algorithm used in generating the Digest value. Permitted values are: - "3" - DSA/SHA-1 - "5" - RSA/SHA-1 - "7" - RSASHA1-NSEC3-SHA1 - "8" - RSA/SHA-256 - "10" - RSA/SHA-512 - "12" - GOST R 34.10-2001 - "13" - ECDSA/SHA-256 - "14" - ECDSA/SHA-384 |
| Digest     | string | Required  | |
| DigestType | string | Required |                                                                                                                                     |
| KeyTag     | string | Required  | |
| ResponseType | string | .Optional | The format that the system respond in. Permitted values are "HTML", "XML" for "TEXT" format. The default response format is "TEXT".                                                                  |

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

AddDnsSec

GetDnsSec