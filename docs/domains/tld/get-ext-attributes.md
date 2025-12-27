GetExtAttributes
================

This command retrieves the extended attributes for a country code TLD \(required parameters specific to the country code\).

Usage
-----

Use this command to determine whether a country code TLD requires extended attributes, and what they are.

Extended attributes supplied by this command are used in the Preconfigure command to configure some TLDs.

Input parameter names for the Preconfigure command are tagged `<Name>` in the return of this command, and permitted values for the Preconfigure command are tagged `<Value>` in the return for this command.

Availability
------------

All resellers have access to this command.

Implementation on eNom.com
---------------------------

This command is not implemented on enom.com.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The country code TLD must be valid.

Input Parameters
----------------

Build the query string using this syntax:

```
https://resellertest.enom.com/interface.asp?
command=nameofcommand&uid=yourloginid&pw=yourpassword
&paramname=paramvalue&nextparamname=nextparamvalue
```
| Input Parameter | Status | Description                      | Max Size |
| --------------- | ----------------------------- | ----------------------------------------------------- | -------- |
| UID       | Required | Account login ID                   | 20 |
| PW | Required           | Account password | 20    |
| TLD       | Required | Top-level domain name \(extension\)         | 15 |
| ResponseType | Optional Format of response. | Permitted values are Text \(default\), HTML, or XML. | 4    |

Returned Parameters and Values
------------------------------

| Output Parameter | Description |
| ---------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ID        | ID number for our internal use |
| Name | Parameter name. Use this as the input parameter name in the Preconfigure command.                                                                                    |
| Value      | Parameter value for the query string. Use this as the input parameter value in the Preconfigure command. |
| Title | Short definition of the parameter value                                                                                                          |
| Application   | Application. 2 indicates Registrant contact. |
| UserDefined | User must supply the value for this parameter from outside sources                                                                                            |
| Required     | Obligation of this parameter. 0 indicates optional, 1 indicates required, 2 indicates a child attribute/parameter that is require d for some values of the parent. Values that require the child are indicated in the `<Required>` node of the parent. |
| Description | Extended definition of the parameter value                                                                                                        |
| IsChild     | If IsChild=1, this parameter is the child of another extended attribute. |
| Command | Name of command executed                                                                                                                 |
| ErrCount     | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values. |
| ErrX | Error messages explaining the failure. These can be presented as is back to the client.                                                                                 |
| Done       | True indicates this entire response has reached you successfully. |

Notes
-----
- The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.
- Check the return parameter ter ErrCount. Err\(ErrCount\) can be presented to the client. Otherwise process the returned parameters as defined above.

Example
-------

The following query requests the extended attributes for the .us TLD, and sends the response in XML, HTML, or Text format.

```
https://resellertest.enom.com/interface.asp?
command=getextattributes&uid=resellid
&pw=resellpw&tld=us&responsetype=xml
```
```
https://resellertest.enom.com/interface.asp?
command=getextattributes&uid=resellid
&pw=resellpw&tld=us&responsetype=html
```
```
https://resellertest.enom.com/interface.asp?
command=getextattributes&uid=resellid
&pw=resellpw&tld=us&responsetype=text
```
The response indicates that the entended attributes for .us are us\_nexus, global\_cc\_us, and us\_purpose. In the Preconfigure command, these parameters would be input, for example, as us\_nexus=C11&us\_purpose=P2:=P2:

```
<?xml version="1.0" ?>
<interface-response>
<Attributes>
 <Attribute>
 <ID>1</ID>
 <Name>us_nexus</Name>
 <Application>2</Application>
 <UserDefined>False</UserDefined>
 <Required>1</Required>
 <Description>Nexus Category</Description>
 <IsChild>0</IsChild>
 <Options>
  <Option>
  <ID>12</ID>
  <Value>C11</Value>
  <Title>US Citizen</Title>
  <Description>A natural person who is a US Citizen</Description>
  </Option>
  <Option>
  <ID>13</ID>
  <Value>C12</Value>
  <Title>Permanent Resident</Title>
  <Description>A natural person who is a Permanent Resident</Description>
  </Option>
.
.
.
 </Options>
 </Attribute>
 <Attribute>
 <ID>4</ID>
 <Name>global_cc_us</Name>
 <Application>2</Application>
 <UserDefined>False</UserDefined>
 <Required>0</Required>
 <Description>Country</Description>
 <IsChild>1</IsChild>
 <Options />
 </Attribute>
 <Attribute>
 <ID>2</ID>
 <Name>us_purpose</Name>
 <Application>2</Application>
 <UserDefined>False</UserDefined>
 <Required>1</Required>
 <Description>Application Purpose</Description>
 <IsChild>0</IsChild>
 <Options>
  <Option>
  <ID>17</ID>
  <Value>P1</Value>
  <Title>For Profit</Title>
  <Description>Business use for profit</Description>
  </Option>
  <Option>
  <ID>18</ID>
  <Value>P2</Value>
  <Title>Non-profit</Title>
  <Description>
Non-profit business, club, association, religious organization, etc.
</Description>
  </Option>
.
.
.
 </Options>
 </Attribute>
</Attributes>
<Command>GETEXTATTRIBUTES</Command>
<ErrCount>0</ErrCount>
<Server>Reseller3</Server>
<Site>enom</Site>
<IsLockable>False</IsLockable>
<IsRealTimeTLD>True</IsRealTimeTLD>
<Done>true</Done>
<debug>
 <![CDATA [ ] ]>
</debug>
</interface-response>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T1
;Encoding Type is utf-8
ID: 1
Name: us_nexus
Application: 2
UserDefined: False
Required: 1
Description: Nexus Category
IsChild: 0
ID: 12
Value: C11
Title: US Citizen
Description: A natural person who is a US Citizen.
ID: 13
Value: C12
Title: Permanent Resident
Description: A natural person who is a Permanent Resident.
ID: 14
Value: C21
Title: Business Entity
Description: An entity or organization that is (i) incorporated within one of the fifty US states, the District of Columbia, or any of the US possessions or territories, or (ii) organized or otherwise constituted under the laws of a state of the US, the District of Columbia or any of its possessions and territories (including federal, state, or local government of the US, or a political subdivision thereof, and non-commercial organizations based in the US.).
.
.
.
Command: GETEXTATTRIBUTES
APIType: API.NET
Language: eng
ErrCount: 0
ResponseCount: 0
MinPeriod: 1
MaxPeriod: 10
Server: sjl0vwresell_t1
Site: eNom
IsLockable:
IsRealTimeTLD:
TimeDifference: +0.00
ExecTime: 0.641
Done: true
TrackingKey: e525d87f-990f-48d1-821e-7a504c1072c2
RequestDateTime: 2/3/2015 6:17:14 PM
```
Related Commands
----------------

AddContact

Contacts

GetAddressBook

GetContacts

GetWhoisContact

Preconfigure