SetHosts
========

Set host records

Usage
-----

Set host records for a domain name.

> To set SRV records, use the ***SetDomainSRVHosts*** command.
>
>
>
> To enable or disable email, use the ***ServiceSelect*** command.

> This command deletes the existing host records before replacing them with the new host records included in the query string. To avoid unpleasant surprises for the user, some resellers use the ***GetHosts*** command to retrieve existing host records, and populate the ***SetHosts*** query string with existing host records while also allowing the user to add or delete host records.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must belong to this account.

Products
--------
- [Google Apps](#vas-add-googleapps)

Input Parameters
----------------

| Parameter           | Type | Status                      | Description |
| ----------------------------- | ------- | ------------------------------------------------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command            | string | Required                     | SetHosts. |
| UID | string | Required | Your Account ID.                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| PW              | string | Required                     | Your API Token. |
| SLD | string | Required | Second-level domain name \(e.g. "enom" in "enom.com"\).                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| TLD              | string | Required                     | Top-level domain name \(extension\). |
| HostNameX X=1 to maximum 50 | integer | Required | Name of the host record to add. This command deletes all existing host records and replaces them with the records you supply in this query string. Index the parameters for each host record X; X must be numeric and in sequence beginning with 1.                                                                                                                                                                                                                                                                                                                                                     |
| RecordTypeX X=1 to maximum 50 | string | Required                     | Record type of host record X. Permitted values: - A -- IP address - AAAA -- IPv6 address - CNAME -- Alias record type, to associate a host name - URL -- URL redirect - FRAME -- Frame redirect - MX -- Mail. Can be a host name under this domain name or the name of a mail - MXE -- Mail Easy \(email forwarding\) - TXT -- Text \(SPF\) record |
| AddressX X=1 to maximum 50 | string | Required | Address to redirect to. - If RecordTypeX=A, AddressX must be an IP address - If RecordTypeX=AAAA, AddressX must be an IPv6 address - If RecordTypeX=CNAME, AddressX must be a fully qualified domain name \(see Note\) or a host name defined in this domain - If RecordTypeX=URL, AddressX must be the exact URL of the page you redirect to, or an IP address, or a fully qualified domain name \(see Note\) - If RecordTypeX=FRAME, AddressX is the actual URL, or the IP address, or the fully qualified domain name \(see Note\) of the page you want to display when someone types Your\_Domain.com - If RecordTypeX=MX, AddressX must be a fully qualified domain name \(see Note\) or a host name defined in this domain - If RecordTypeX=MXE, AddressX must be an IP address - If RecordTypeX=TXT, AddressX is a text \(SPF\) record. For help writing an SPF record, go to [http://spf.pobox.com/wizard.html](http://spf.pobox.com/wizard.md) |
| MXPrefX X=1 to maximum 50\)  | string | Optional; use with record type MX; default is 10 | Host record preference for setting mail redirection. The lower the number, the higher the priority |
| ResponseType | string | Optional | Format of response. Permitted values are: Text \(default\), HTML, or XML.                                                                                                                                                                                                                                                                                                                                                                                                                                           |

Returned Parameters and Values
------------------------------

> Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Parameter | Type | Description                                           |
| ---------- | ------ | ------------------------------------------------------------------------------------------------ |
| Command  | string | Name of command executed.                                    |
| ErrorCount | int | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX    | string | Error messages explaining the failure. These can be presented as is back to the client.    |
| Done    | string | "True" value indicates this entire response has reached you successfully.            |

Notes
-----
- A fully qualified domain name is expressed in the format “hostname.SLD.TLD.”. Note that the period at the end is an essential component of a fully qualified domain name.
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0, the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, process the returned parameters as defined above.

Examples
--------

```
https://resellertest.enom.com/interface.asp?command=SetHosts&uid=YourAccountID&pw=YourApiToken&SLD={Required}&TLD={Required}&HostName1=@&RecordType1=A
&Address1=66.150.5.189&HostName2=photos&RecordType2=CNAME
&Address2=photos.msn.com.&HostName3=yahoo&RecordType3=URL
&Address3=204.71.200.72&HostName4=msn&RecordType4=FRAME
&Address4=http://www.msn.com&responsetype=xml
```
```
<?xml version="1.0" ?>
<interface-response>
 <DomainRRP>E</DomainRRP>
 <Command>SETHOSTS</Command>
 <ErrCount>0</ErrCount>
 <Server>sjl2vwapi01</Server>
 <Site>enom</Site>
 <IsLockable>True</IsLockable>
 <IsRealTimeTLD>True</IsRealTimeTLD>
 <Done>true</Done>
</interface-response>
```
Related Commands
----------------

GetHosts

GetMetaTag

GetRegHosts

GetSPFHosts

[PurchaseServices](../docs/purchaseservices.md)

[SetDNSHost](../docs/setdnshost.md)

SetDomainSRVHosts

SetSPFHosts

UpdateMetaTag