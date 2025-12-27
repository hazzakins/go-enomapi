GetRegHosts
===========

Get host records for a domain without mail host records.

Usage
-----

Use this command to display host record information for a domain. The response does not include mail host entries. GetHosts, a similar command, returns all host records including mail host records.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

Log on to resellertest.enom.com with Login ID resellid, password resellpw.

[https://resellertest.enom.com/domains/DomainDetail.asp?domainnameid=152533676](https://resellertest.enom.com/domains/DomainDetail.asp?domainnameid=152533676)

In the domain control panel, DNS settings box, the configure button calls the GetRegHosts command. GetRegHosts command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The domain name must belong to this account.
- The domain must use eNom’s domain name servers.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                                           | Max Size |
| --------------- | ----------------------------- | ------------------------------------------------------------------------------------------------ | -------- |
| UID       | Required | Account login ID                                         | 20 |
| PW | Required           | Account password | 20    |
| SLD       | Required | Second-level domain name \(for example, enom in enom.com\)                   | 63 |
| TLD | Required           | Top-level domain name \(extension\) | 15    |
| ExtFormat    | Optional | ExtFormat=1 encloses the repeating host information in a separate tag so it is easier to parse. | 1 |
| ResponseType | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | ------------------------------------------------------------------------------------------------ |
| HostIDX     | Identification number of this individual host record |
| HostNameX | Name of this individual host record                               |
| AddressX     | Address of this individual host record |
| RecordTypeX | Record type of this individual host record                            |
| Command     | Name of command executed |
| ErrCount | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX       | Error messages explaining the failure. These can be presented as is back to the client. |
| Done | True indicates this entire response has reached you successfully.                |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query requests the DNS host records for a domain name, specifies ExtFormat=1 to enclose the host records in a separate tag, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=getreghosts&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&ExtFormat=1&responsetype=XML
```
```
https://resellertest.enom.com/interface.asp?
command=getreghosts&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&ExtFormat=1&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=getreghosts&uid=resellid&pw=resellpw
&sld=resellerdocs&tld=com&ExtFormat=1&responsetype=text
```
The response is as follows:

```
<interface-response>
 <hostrecords>
 <host>
  <hostID/>
  <HostName>www</HostName>
  <Address/>
  <RecordType>CNAME</RecordType>
 </host>
 <host>
  <hostID>18149735</hostID>
  <HostName>@ (none)</HostName>
  <Address>66.150.5.189</Address>
  <RecordType>A</RecordType>
 </host>
 <host>
  <hostID>18487896</hostID>
  <HostName>msn</HostName>
  <Address>http://www.msn.com</Address>
  <RecordType>FRAME</RecordType>
 </host>
 <host>
  <hostID>18487894</hostID>
  <HostName>photos</HostName>
  <Address>photos.msn.com.</Address>
  <RecordType>CNAME</RecordType>
 </host>
 <host>
  <hostID>18487895</hostID>
  <HostName>yahoo</HostName>
  <Address>204.71.200.72</Address>
  <RecordType>URL</RecordType>
 </host>
 <host>
  <hostID/>
  <HostName/>
  <Address/>
  <RecordType>A</RecordType>
 </host>
 <host>
  <hostID/>
  <HostName/>
  <Address/>
  <RecordType>A</RecordType>
 </host>
 <host>
  <hostID/>
  <HostName/>
  <Address/>
  <RecordType>A</RecordType>
 </host>
 <host>
  <hostID/>
  <HostName/>
  <Address/>
  <RecordType>A</RecordType>
 </host>
 <host>
  <hostID/>
  <HostName>* (other)</HostName>
  <Address/>
  <RecordType>CNAME</RecordType>
 </host>
 </hostrecords>
 <ParkingEnabled>False</ParkingEnabled>
 <ServiceSelect>1114</ServiceSelect>
 <Command>GETREGHOSTS</Command>
 <Language>eng</Language>
 <ErrCount>0</ErrCount>
 <ResponseCount>0</ResponseCount>
 <MinPeriod>1</MinPeriod>
 <MaxPeriod>10</MaxPeriod>
 <Server>SJL21WRESELLT01</Server>
 <Site>eNom</Site>
 <IsLockable>True</IsLockable>
 <IsRealTimeTLD>True</IsRealTimeTLD>
 <TimeDifference>+08.00</TimeDifference>
 <ExecTime>1.047</ExecTime>
 <Done>true</Done>
 <RequestDateTime>12/9/2011 2:26:00 AM</RequestDateTime>
 <debug></debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
hostID1:
HostName1: www
Address1:
RecordType1: CNAME
hostID2: 19920062
HostName2: @ (none)
Address2: 66.150.5.189
RecordType2: A
hostID3: 19995603
HostName3: msn
Address3: http://www.msn.com
RecordType3: FRAME
hostID4: 19987522
HostName4: photos
Address4: photos.msn.com.
RecordType4: CNAME
hostID5: 19987523
HostName5: yahoo
Address5: 204.71.200.72
RecordType5: URL
hostID6:
HostName6:
Address6:
RecordType6: A
hostID7:
HostName7:
Address7:
RecordType7: A
hostID8:
HostName8:
Address8:
RecordType8: A
hostID9:
HostName9:
Address9:
RecordType9: A
hostID10:
HostName10: * (other)
Address10:
RecordType10: CNAME
HostCount: 10
ParkingEnabled: False
ServiceSelect: 1048
Command: GETREGHOSTS
APIType: API
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: SJL0VWRESELL_T
Site: eNom
IsLockable: True
IsRealTimeTLD: True
TimeDifference: +08.00
ExecTime: 0.203
Done: true
RequestDateTime: 2/4/2015 11:31:35 AM
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
hostID1=
HostName1=www
Address1=
RecordType1=CNAME
hostID2=19920062
HostName2=@ (none)
Address2=66.150.5.189
RecordType2=A
hostID3=19995603
HostName3=msn
Address3=http://www.msn.com
RecordType3=FRAME
hostID4=19987522
HostName4=photos
Address4=photos.msn.com.
RecordType4=CNAME
hostID5=19987523
HostName5=yahoo
Address5=204.71.200.72
RecordType5=URL
hostID6=
HostName6=
Address6=
RecordType6=A
hostID7=
HostName7=
Address7=
RecordType7=A
hostID8=
HostName8=
Address8=
RecordType8=A
hostID9=
HostName9=
Address9=
RecordType9=A
hostID10=
HostName10=* (other)
Address10=
RecordType10=CNAME
HostCount=10
ParkingEnabled=False
ServiceSelect=1048
Command=GETREGHOSTS
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=1
MaxPeriod=10
Server=SJL0VWRESELL_T1
Site=eNom
IsLockable=True
IsRealTimeTLD=True
TimeDifference=+08.00
ExecTime=0.109
Done=true
RequestDateTime=2/4/2015 11:32:02 AM
```
Related Commands
----------------

GetHosts

GetMetaTag

SetDNSHost

SetHosts

UpdateMetaTag