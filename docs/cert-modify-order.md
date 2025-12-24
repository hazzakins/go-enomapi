CertModifyOrder
===============

Cancel a certificate request

Usage
-----

Cancel a cert configuration order.

> ### Cancellation is permitted while a cert is in status Approval email sent or Processing. This command does not delete the cert from a customer’s account; it only cancels the configuration order. This allows a customer to change the configuration, including the domain name with which the cert will be associated.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The cert must belong to this account.

Input Parameters
----------------

| Input Parameter | Type | Status  | Description |
| --------------- | ------ | -------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| Command     | string | Required | CertModifyOrder. |
| UID | string | Required | Your Account ID.                                                           |
| PW       | string | Required | Your API Token. |
| CertID | string | Required | ID number of this cert. Use the [CertGetCerts](../docs/cert-get-certs.md) command to retrieve the ID number.            |
| ResponseType  | string | Optional | The format that the system respond in. Permitted values are "HTML", "XML" for "TEXT" format. The default response format is "TEXT". |

Returned Parameters and Values
------------------------------

Check the return parameter "ErrCount". If greater than 0 the transaction failed. The parameter "Err\(ErrCount\)" can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter | Type  | Description |
| ---------------- | ------- | ------------------------------------------------------------------------------------------------ |
| Command | string | CertModifyOrder. |
| CertID | boolean | Identification number of this cert. |
| ErrCount | int   | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | string | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | string | "True" indicates this entire response has reached you successfully. |

Example Input / Output
----------------------

```
https://resellertest.enom.com/interface.aspcommand=certmodifyorder&uid=resellid&pw=resellpw&CertID=48455&ResponseType=XML
```
```
<?xml version="1.0" ?>
<interface-response>
 <CertModifyOrder>
  <CertID>48455</CertID>
 </CertModifyOrder>
 <Command>CERTMODIFYORDER</Command>
 <Language>en</Language>
 <ErrCount>0</ErrCount>
 <ResponseCount>0</ResponseCount>
 <MinPeriod />
 <MaxPeriod>10</MaxPeriod>
 <Server>RESELLERTEST</Server>
 <Site>enom</Site>
 <IsLockable />
 <IsRealTimeTLD />
 <TimeDifference>+0.00</TimeDifference>
 <ExecTime>0.109375</ExecTime>
 <Done>true</Done>
</interface-response>
```
Related Commands
----------------

[CertChangeApproverEmail](../docs/cert-change-approveremail.md)

[CertConfigureCert](../docs/cert-configure-cert.md)

[CertGetApproverEmail](../docs/cert-get-approver-email.md)

[CertGetCertDetail](../docs/cert-get-certdetail.md)

[CertGetCerts](../docs/cert-get-certs.md)

[CertModifyOrder](../docs/cert-modify-order.md)

[CertParseCSR](../docs/cert-parse-csr.md)

[CertPurchaseCert](../docs/cert-purchase-cert.md)

[CertResendApproverEmail](../docs/cert-resend-approveremail.md)