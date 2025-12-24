CheckLogin
==========

Validate user login for a domain name.

Usage
-----

Use this command when you want to validate a user’s identity.

Availability
------------

All resellers have access to this command.

Constraints
-----------

The query must meet the following requirements:
- The login ID and password must be valid.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=CheckLogin&uid={YourAccountID}&pw={YourApiToken}&responsetype={Optional}
```
| Input Parameter | Type | Description                                      |
| --------------- | ------ | -------------------------------------------------------------------------------------- |
| command     | string | CheckLogin                                       |
| uid       | string | Your Account ID                                    |
| pw       | string | Your API Token                                     |
| responsetype  | string | Optional - Format of response. Permitted values are Text \(default\), HTML, or XML. |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0, the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Output Parameter   | Type | Description                                                |
| --------------------- | ------ | --------------------------------------------------------------------------------------------------------- |
| command        | string | Name of command executed                                         |
| ErrCount       | int | The number of errors if any occurred. If greater than 0, check the Err\(1 to ErrCount\) values.     |
| err\{X\}       | string | Error messages explaining the failure. These can be presented as-is back to the client.        |
| done         | string | "True" value indicates this entire response has reached you successfully.                |
| PartyID        | int | PartyID of the account                                          |
| Reseller       | string | Is this a reseller account? True False                                  |
| ParentAccount     | string | Parent account ID number. A returned value of "000-00-0000" means this is a direct account with eNom. |
| IsETP         | int | Is this account an ETP account.                                     |
| SEOPartner      | string | Is this an SEO partner.                                         |
| BulkMembershipType  | string | BulkRegister membership type                                       |
| BulkMembership    | string | Is this account a BulkRegister member?                                  |
| bBulkMembershipStatus | int | Status of BulkRegister membership                                     |
| VerifiedBidder    | string | Is this a NameJet verified bidder.                                    |
| RSA          | string | Reserved parameter. Should always return False.                             |
| SiteType       | string | Site                                                   |
| HelpEmailAddress   | string | Email address for the Billing contact for this account                          |
| Type         | int | Agreement type                                              |
| Description      | string | Name of this agreement                                          |
| Category       | string | Category of this agreement                                        |
| RemoteURL       | string | URL of this agreement, if it is not served by us                             |
| LatestVersion     | float | Version number of this agreement                                     |
| Agreed        | int | Has this account agreed to this agreement.                                |
| CurrentCompliance   | int | Is this account currently complying with this agreement.                         |
| AgreementNote     | string | Note on this agreement                                          |
| Enforced       | string | Enforcement setting                                            |
| ViewAgreement     | int | View setting for agreement                                        |
| ClubDrop       | string | Is this account a Club Drop member?                                    |
| ClubDropExpired    | string | Is this Club Drop membership expired                                   |
| CartItems       | int | Number of items currently in cart                                     |
| ResponseCount     | int | Number of responses                                            |
| MinPeriod       | int | Minimum registration period \(for domain names\)                             |
| MaxPeriod       | int | Maximum registration period \(for domain names\)                             |
| Server        | string | Server name that executed the API command.                                |
| Site         | string | Site for which this response is configured                                |
| IsLockable      | string | Can this domain be Registrar locked. \(for domain names\)                        |
| IsRealTimeTLD     | string | Is this a domain that can be purchased in real time. \(for domain names\)                |
| TimeDifference    | string | Time difference between this site and our servers                             |
| ExecTime       | string | Time elapsed to execute this command                                   |
| pin          | int | PIN number for direct eNom accounts for priority phone access.                      |

Notes
-----

The default response format is plain text. To receive the response in HTML or XML format, set the "ResponseType=" flag to either "HTML" or "XML" in your request.

Check the return parameter "ErrCount". If it is greater than 0, the transaction has failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, process the returned parameters as defined above.

Example of Output
-----------------

The following query requests the party ID and reseller status of account resellid and requests the response in format of either "text" \(default\), "xml" or "html"

```
<interface-response>
<PartyID>{BB4A2DE1-6485-45CB-A4FC-EE39BA0C1629}</PartyID>
<Reseller>True</Reseller>
<ParentAccount>000-00-0000</ParentAccount>
<IsETP>1</IsETP>
<SEOPartner>NA</SEOPartner>
<BulkMembershipType>None</BulkMembershipType>
<BulkMembership>False</BulkMembership>
<BulkMembershipStatus/>
<bBulkMembershipStatus>0</bBulkMembershipStatus>
<VerifiedBidder>False</VerifiedBidder>
<RSA>False</RSA>
<SiteType>E</SiteType>
<HelpEmailAddress>[email protected]</HelpEmailAddress>
<AgreementStatus>
<Agreement>
<Type>1</Type>
<Description>ResellerAgreement</Description>
<Category>ResellerAgreement</Category>
<RemoteURL/>
<LatestVersion>2.5</LatestVersion>
<Agreed>1</Agreed>
<CurrentCompliance>1</CurrentCompliance>
<AgreementNote/>
</Agreement>
<Agreement>
<Type>2</Type>
<Description>Network Solutions Service Agreement</Description>
<Category>ClubDrop</Category>
<RemoteURL>http://www.networksolutions.com/legal/static-service-agreement.jsp</RemoteURL>
<LatestVersion>1.0</LatestVersion>
<Agreed>0</Agreed>
<CurrentCompliance>0</CurrentCompliance>
<AgreementNote/>
</Agreement>
<Agreement>
<Type>3</Type>
<Description>NameJet Terms of Use</Description>
<Category>ClubDrop</Category>
<RemoteURL/>
<LatestVersion>2.1</LatestVersion>
<Agreed>0</Agreed>
<CurrentCompliance>0</CurrentCompliance>
<AgreementNote/>
</Agreement>
<Agreement>
<Type>4</Type>
<Description>ResellerNotification</Description>
<Category>ResellerNotification</Category>
<RemoteURL/>
<LatestVersion>1.3</LatestVersion>
<Agreed>1</Agreed>
<CurrentCompliance>1</CurrentCompliance>
<AgreementNote/>
</Agreement>
<Agreement>
<Type>5</Type>
<Description>eNomCentral Registration Agreement</Description>
<Category>ClubDrop</Category>
<RemoteURL>http://www.enomcentral.com/terms/agreement.asp</RemoteURL>
<LatestVersion>10.0</LatestVersion>
<Agreed>0</Agreed>
<CurrentCompliance>1</CurrentCompliance>
<AgreementNote/>
</Agreement>
<Agreement>
<Type>6</Type>
<Description>RegistrationAgreement</Description>
<Category>RegistrationAgreement</Category>
<RemoteURL/>
<LatestVersion>2.0</LatestVersion>
<Agreed>1</Agreed>
<CurrentCompliance>1</CurrentCompliance>
<AgreementNote/>
</Agreement>
<Agreement>
<Type>7</Type>
<Description>BulkRegister temp</Description>
<Category>Admin</Category>
<RemoteURL/>
<LatestVersion>1.0</LatestVersion>
<Agreed>0</Agreed>
<CurrentCompliance>1</CurrentCompliance>
<AgreementNote/>
</Agreement>
</AgreementStatus>
<AgreementCount/>
<Account>217-no-0647</Account>
<IsETP>1</IsETP>
<SecurePackageEnabled>False</SecurePackageEnabled>
<ForceAddedSecurity>False</ForceAddedSecurity>
<SecureLockToken/>
<FailureCount>0</FailureCount>
<SecureLockReturnMessage/>
<ExtraCheck>False</ExtraCheck>
<ViewAgreement>False</ViewAgreement>
<ClubDrop>True</ClubDrop>
<ClubDropExpired>False</ClubDropExpired>
<CartItems>3</CartItems>
<PdqAccountId>1000186</PdqAccountId>
<RSFAccountStatus>New</RSFAccountStatus>
<RSFExpDate>7/12/2012</RSFExpDate>
<IsBadEmail>False</IsBadEmail>
<Command>CHECKLOGIN</Command>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod/>
<MaxPeriod>10</MaxPeriod>
<Server>SJL21WRESELLT01</Server>
<Site>eNom</Site>
<IsLockable/>
<IsRealTimeTLD/>
<TimeDifference>+0.00</TimeDifference>
<ExecTime>0.608</ExecTime>
<Done>true</Done>
<RequestDateTime>12/7/2011 4:15:00 AM</RequestDateTime>
<debug></debug>
</interface-response>
```
```
;URL Interface<br>
;Machine is SJL0VWRESELL_T<br>
;Encoding Type is utf-8<br>
<HTML><BODY>
<STRONG>PartyID= </STRONG>{BB4A2DE1-6485-45CB-A4FC-EE39BA0C1629}<BR>
<STRONG>Reseller= </STRONG>True<BR>
<STRONG>ParentAccount= </STRONG>000-00-0000<BR>
<STRONG>IsETP= </STRONG>1<BR>
<STRONG>SEOPartner= </STRONG>NA<BR>
<STRONG>eNomRetailGroup= </STRONG>False<BR>
<STRONG>BulkMembershipType= </STRONG>None<BR>
<STRONG>BulkMembership= </STRONG>False<BR>
<STRONG>BulkMembershipStatus= </STRONG><BR>
<STRONG>bBulkMembershipStatus= </STRONG>0<BR>
<STRONG>VerifiedBidder= </STRONG>False<BR>
<STRONG>RSA= </STRONG>False<BR>
<STRONG>SiteType= </STRONG>E<BR>
<STRONG>HelpEmailAddress= </STRONG>[email protected]<BR>
<STRONG>Type1= </STRONG>1<BR>
<STRONG>Description1= </STRONG>ResellerAgreement<BR>
<STRONG>Category1= </STRONG>ResellerAgreement<BR>
<STRONG>RemoteURL1= </STRONG><BR>
<STRONG>LatestVersion1= </STRONG>3.0<BR>
<STRONG>Agreed1= </STRONG>1<BR>
<STRONG>CurrentCompliance1= </STRONG>1<BR>
<STRONG>AgreementNote1= </STRONG><BR>
<STRONG>Type2= </STRONG>2<BR>
<STRONG>Description2= </STRONG>Network Solutions Service Agreement<BR>
<STRONG>Category2= </STRONG>ClubDrop<BR>
<STRONG>RemoteURL2= </STRONG>http://www.networksolutions.com/legal/static-service-agreement.jsp<BR>
<STRONG>LatestVersion2= </STRONG>1.0<BR>
<STRONG>Agreed2= </STRONG>0<BR>
<STRONG>CurrentCompliance2= </STRONG>0<BR>
<STRONG>AgreementNote2= </STRONG><BR>
<STRONG>Type3= </STRONG>3<BR>
<STRONG>Description3= </STRONG>NameJet Terms of Use<BR>
<STRONG>Category3= </STRONG>ClubDrop<BR>
<STRONG>RemoteURL3= </STRONG><BR>
<STRONG>LatestVersion3= </STRONG>2.1<BR>
<STRONG>Agreed3= </STRONG>0<BR>
<STRONG>CurrentCompliance3= </STRONG>0<BR>
<STRONG>AgreementNote3= </STRONG><BR>
<STRONG>Type4= </STRONG>4<BR>
<STRONG>Description4= </STRONG>ResellerNotification<BR>
<STRONG>Category4= </STRONG>ResellerNotification<BR>
<STRONG>RemoteURL4= </STRONG><BR>
<STRONG>LatestVersion4= </STRONG>3.2<BR>
<STRONG>Agreed4= </STRONG>1<BR>
<STRONG>CurrentCompliance4= </STRONG>1<BR>
<STRONG>AgreementNote4= </STRONG><BR>
<STRONG>Type5= </STRONG>5<BR>
<STRONG>Description5= </STRONG>eNomCentral Registration Agreement<BR>
<STRONG>Category5= </STRONG>ClubDrop<BR>
<STRONG>RemoteURL5= </STRONG>http://www.enomcentral.com/terms/agreement.asp<BR>
<STRONG>LatestVersion5= </STRONG>10.0<BR>
<STRONG>Agreed5= </STRONG>0<BR>
<STRONG>CurrentCompliance5= </STRONG>1<BR>
<STRONG>AgreementNote5= </STRONG><BR>
<STRONG>Type6= </STRONG>6<BR>
<STRONG>Description6= </STRONG>RegistrationAgreement<BR>
<STRONG>Category6= </STRONG>RegistrationAgreement<BR>
<STRONG>RemoteURL6= </STRONG><BR>
<STRONG>LatestVersion6= </STRONG>2.0<BR>
<STRONG>Agreed6= </STRONG>1<BR>
<STRONG>CurrentCompliance6= </STRONG>1<BR>
<STRONG>AgreementNote6= </STRONG><BR>
<STRONG>Type7= </STRONG>7<BR>
<STRONG>Description7= </STRONG>BulkRegister temp<BR>
<STRONG>Category7= </STRONG>Admin<BR>
<STRONG>RemoteURL7= </STRONG><BR>
<STRONG>LatestVersion7= </STRONG>1.0<BR>
<STRONG>Agreed7= </STRONG>0<BR>
<STRONG>CurrentCompliance7= </STRONG>1<BR>
<STRONG>AgreementNote7= </STRONG><BR>
<STRONG>Type8= </STRONG>99<BR>
<STRONG>Description8= </STRONG>MultipleNotices<BR>
<STRONG>Category8= </STRONG>MultipleNotices<BR>
<STRONG>RemoteURL8= </STRONG><BR>
<STRONG>LatestVersion8= </STRONG>2.0<BR>
<STRONG>Agreed8= </STRONG>1<BR>
<STRONG>CurrentCompliance8= </STRONG>1<BR>
<STRONG>AgreementNote8= </STRONG><BR>
<STRONG>AgreementCount= </STRONG>9<BR>
<STRONG>Account= </STRONG>217-no-0647<BR>
<STRONG>SecurePackageEnabled= </STRONG>False<BR>
<STRONG>ForceAddedSecurity= </STRONG>False<BR>
<STRONG>SecureLockToken= </STRONG><BR>
<STRONG>FailureCount= </STRONG>0<BR>
<STRONG>SecureLockReturnMessage= </STRONG><BR>
<STRONG>ViewAgreement= </STRONG>False<BR>
<STRONG>ClubDrop= </STRONG>True<BR>
<STRONG>ClubDropExpired= </STRONG>False<BR>
<STRONG>CartItems= </STRONG>34<BR>
<STRONG>PdqAccountId= </STRONG>1000186<BR>
<STRONG>RSFAccountStatus= </STRONG>Active<BR>
<STRONG>RSFExpDate= </STRONG>7/12/2015<BR>
<STRONG>IsBadEmail= </STRONG>False<BR>
<STRONG>pin= </STRONG>2619237<BR>
<STRONG>enabled= </STRONG>True<BR>
<STRONG>Command= </STRONG>CHECKLOGIN<BR>
<STRONG>APIType= </STRONG>API<BR>
<STRONG>Language= </STRONG>eng<BR>
<STRONG>ErrCount= </STRONG>0<BR>
<STRONG>ResponseCount= </STRONG>0<BR>
<STRONG>MinPeriod= </STRONG>1<BR>
<STRONG>MaxPeriod= </STRONG>10<BR>
<STRONG>Server= </STRONG>SJL0VWRESELL_T<BR>
<STRONG>Site= </STRONG>eNom<BR>
<STRONG>IsLockable= </STRONG><BR>
<STRONG>IsRealTimeTLD= </STRONG><BR>
<STRONG>TimeDifference= </STRONG>+0.00<BR>
<STRONG>ExecTime= </STRONG>0.188<BR>
<STRONG>Done= </STRONG>true<BR>
<STRONG>RequestDateTime= </STRONG>2/3/2015 2:39:03 PM<BR>
</BODY></HTML>
```
```
;URL Interface
;Machine is SJL0VWRESELL_T
;Encoding Type is utf-8
PartyID={BB4A2DE1-6485-45CB-A4FC-EE39BA0C1629}
Reseller=True
ParentAccount=000-00-0000
IsETP=1
SEOPartner=NA
eNomRetailGroup=False
BulkMembershipType=None
BulkMembership=False
BulkMembershipStatus=
bBulkMembershipStatus=0
VerifiedBidder=False
RSA=False
SiteType=E
[email protected]
Type1=1
Description1=ResellerAgreement
Category1=ResellerAgreement
RemoteURL1=
LatestVersion1=3.0
Agreed1=1
CurrentCompliance1=1
AgreementNote1=
Type2=2
Description2=Network Solutions Service Agreement
Category2=ClubDrop
RemoteURL2=http://www.networksolutions.com/legal/static-service-agreement.jsp
LatestVersion2=1.0
Agreed2=0
CurrentCompliance2=0
AgreementNote2=
Type3=3
Description3=NameJet Terms of Use
Category3=ClubDrop
RemoteURL3=
LatestVersion3=2.1
Agreed3=0
CurrentCompliance3=0
AgreementNote3=
Type4=4
Description4=ResellerNotification
Category4=ResellerNotification
RemoteURL4=
LatestVersion4=3.2
Agreed4=1
CurrentCompliance4=1
AgreementNote4=
Type5=5
Description5=eNomCentral Registration Agreement
Category5=ClubDrop
RemoteURL5=http://www.enomcentral.com/terms/agreement.asp
LatestVersion5=10.0
Agreed5=0
CurrentCompliance5=1
AgreementNote5=
Type6=6
Description6=RegistrationAgreement
Category6=RegistrationAgreement
RemoteURL6=
LatestVersion6=2.0
Agreed6=1
CurrentCompliance6=1
AgreementNote6=
Type7=7
Description7=BulkRegister temp
Category7=Admin
RemoteURL7=
LatestVersion7=1.0
Agreed7=0
CurrentCompliance7=1
AgreementNote7=
Type8=99
Description8=MultipleNotices
Category8=MultipleNotices
RemoteURL8=
LatestVersion8=2.0
Agreed8=1
CurrentCompliance8=1
AgreementNote8=
AgreementCount=9
Account=217-no-0647
SecurePackageEnabled=False
ForceAddedSecurity=False
SecureLockToken=
FailureCount=0
SecureLockReturnMessage=
ViewAgreement=False
ClubDrop=True
ClubDropExpired=False
CartItems=34
PdqAccountId=1000186
RSFAccountStatus=Active
RSFExpDate=7/12/2015
IsBadEmail=False
pin=2619237
enabled=True
Command=CHECKLOGIN
APIType=API
Language=eng
ErrCount=0
ResponseCount=0
MinPeriod=
MaxPeriod=10
Server=SJL0VWRESELL_T
Site=eNom
IsLockable=
IsRealTimeTLD=
TimeDifference=+0.00
ExecTime=0.188
Done=true
RequestDateTime=2/3/2015 2:39:03 PM
```
Related Commands
----------------

CreateAccount

CreateSubAccount

GetAccountInfo

GetAllAccountInfo

GetOrderDetail

GetOrderList

GetReport

GetSubAccountDetails

GetSubAccounts

GetTransHistory

UpdateAccountInfo

UpdateCusPreferences