API Error Codes
===============

When you run a query against the eNom API, the response includes a numeric code and a corresponding text message. This allows us to standardize our responses and make them more specific; it also allows our international resellers to translate API responses to their local language.

> ### The numerical response codes are six-digit numbers that indicate the system, type of error, and parameter that are the subjects of this response.
>
>
>
> - The first digit in each six-digit error code indicates the system that is generating the message. For example, error codes that begin with 4 indicate authentication errors. This first digit is referred to as the system bit.
> - The second and third digits in an error code indicate the nature of the error. For example, x04xxx indicates invalid data. These digits are referred to as the error bit.
> - The last three digits in an error code indicate the parameter that is failing. For example, xxx156 indicates a problem with the password. These digits are referred to as the parameter bit.

Error Code Components
---------------------

System bit
----------

*The first digit in each six-digit error code indicates the system that is generating the message. This first digit is referred to as the system bit.*

| System Bit | Definition |
| ---------- | ------------------------------ |
| 0xxxxx   | An unknown error has occurred |
| 1xxxxx | Command completed successfully |
| 2xxxxx   | Registry error |
| 3xxxxx | Validation error        |
| 4xxxxx   | Authentication error |
| 5xxxxx | Payment error         |
| 6xxxxx   | System error |
| 7xxxxx | Policy error          |

Error bit
---------

*The second and third digits in an error code indicate the nature of the error. These digits are referred to as the error bit.*

| Error Bit | Definition |
| --------- | ----------------------------- |
| x00xxx  | Does not apply |
| x01xxx | Missing            |
| x02xxx  | Duplicate |
| x03xxx | Out of range         |
| x04xxx  | Invalid |
| x05xxx | Type mismatch         |
| x06xxx  | Policy violation |
| x07xxx | Connection failed       |
| x08xxx  | Failed to retrieve |
| x09xxx | Failed to add         |
| x10xxx  | Failed to update |
| x11xxx | Failed to delete       |
| x12xxx  | Disabled |
| x13xxx | Unauthorized         |
| x14xxx  | Declined |
| x15xxx | Not available         |
| x16xxx  | Not found |
| x17xxx | Expired            |
| x18xxx  | Down for maintenance |
| x19xxx | Failed to deduct       |
| x20xxx  | Failed to send |
| x21xxx | Failed to process       |
| x22xxx  | Failed to downgrade |
| x23xxx | Failed to upgrade       |
| x24xxx  | Insufficient Reseller balance |
| x99xxx | Unknown            |

Parameter bit
-------------

*The last three digits in an error code indicate the parameter that is failing. These digits are referred to as the parameter bit.*

| Parameter bit | Definition |
| ------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| xxx000    | Does not apply |
| xxx001 | First name \(generic\); for specific contact types see xxx019 through xxx024                                                                       |
| xxx002    | Last name \(generic\); for specific contact types see xxx025 through xxx029 |
| xxx003 | Address \(generic\); for specific contact types see xxx030 through xxx034                                                                        |
| xxx004    | City \(generic\); for specific contact types see xxx040 through xxx044 |
| xxx005 | State/Province \(generic\); for specific contact types see xxx045 through xxx049                                                                     |
| xxx006    | Postal code \(generic\); for specific contact types see xxx050 through xxx054 |
| xxx007 | Country \(generic\); for specific contact types see xxx055 through xxx059                                                                        |
| xxx008    | Job title \(generic\); for specific contact types see xxx060 through xxx064 |
| xxx009 | Organization \(generic\)                                                                                                 |
| xxx010    | Phone number \(generic\); for specific contact types see xxx065 through xxx069 |
| xxx011 | Fax number \(generic\); for specific contact types see xxx070 through xxx074                                                                       |
| xxx012    | Email \(generic\); for specific contact types see xxx075 through xxx079 |
| xxx019 | Registrant contact first name                                                                                              |
| xxx021    | Auxiliary Billing contact first name |
| xxx022 | Technical contact first name                                                                                               |
| xxx023    | Administrative contact first name |
| xxx024 | Billing contact first name                                                                                                |
| xxx025    | Registrant contact last name |
| xxx026 | Auxiliary Billing contact last name                                                                                           |
| xxx027    | Technical contact last name |
| xxx028 | Administrative contact last name                                                                                             |
| xxx029    | Billing contact last name |
| xxx030 | Registrant contact address, line 1                                                                                            |
| xxx031    | Auxiliary Billing contact address, line 1 |
| xxx032 | Technical contact address, line 1                                                                                            |
| xxx033    | Administrative contact address, line 1 |
| xxx034 | Billing contact address, line 1                                                                                             |
| xxx035    | Registrant contact address, line 2 |
| xxx036 | Auxiliary Billing contact address, line 2                                                                                        |
| xxx037    | Technical contact address, line 2 |
| xxx038 | Administrative contact address, line 2                                                                                          |
| xxx039    | Billing contact address, line 2 |
| xxx040 | Registrant contact city                                                                                                 |
| xxx041    | Auxiliary Billing contact city |
| xxx042 | Technical contact city                                                                                                  |
| xxx043    | Administrative contact city |
| xxx044 | Billing contact city                                                                                                   |
| xxx045    | Registrant contact state/province |
| xxx046 | Auxiliary Billing contact state/province                                                                                         |
| xxx047    | Technical contact state/province |
| xxx048 | Administrative contact state/province                                                                                          |
| xxx049    | Billing contact state/province |
| xxx050 | Registrant contact postal code                                                                                              |
| xxx051    | Auxiliary Billing contact postal code |
| xxx052 | Technical contact postal code                                                                                              |
| xxx053    | Administrative contact postal code |
| xxx054 | Billing contact postal code                                                                                               |
| xxx055    | Registrant contact country |
| xxx056 | Auxiliary Billing contact country                                                                                            |
| xxx057    | Technical contact country |
| xxx058 | Administrative contact country                                                                                              |
| xxx059    | Billing contact country |
| xxx060 | Registrant contact job title                                                                                               |
| xxx061    | Auxiliary Billing contact job title |
| xxx062 | Technical contact job title                                                                                               |
| xxx063    | Administrative contact job title |
| xxx064 | Billing contact job title                                                                                                |
| xxx065    | Registrant contact phone |
| xxx066 | Auxiliary Billing contact phone                                                                                             |
| xxx067    | Technical contact phone |
| xxx068 | Administrative contact phone                                                                                               |
| xxx069    | Billing contact phone |
| xxx070 | Registrant contact fax                                                                                                  |
| xxx071    | Auxiliary Billing contact fax |
| xxx072 | Technical contact fax                                                                                                  |
| xxx073    | Administrative contact fax |
| xxx074 | Billing contact fax                                                                                                   |
| xxx075    | Registrant contact email |
| xxx076 | Auxiliary Billing contact email                                                                                             |
| xxx077    | Technical contact email |
| xxx078 | Administrative contact email                                                                                               |
| xxx079    | Billing contact email |
| xxx144 | Party ID                                                                                                         |
| xxx145    | Registrant contact \(use this when you are unable to retrieve the entire block of registrant contact information\) |
| xxx146 | Auxiliary Billing contact \(entire block\)                                                                                        |
| xxx147    | Technical contact \(entire block\) |
| xxx148 | Administrative contact \(entire block\)                                                                                         |
| xxx149    | Billing contact \(entire block\) |
| xxx150 | Command                                                                                                         |
| xxx151    | SLD |
| xxx152 | TLD                                                                                                           |
| xxx153    | Domain name |
| xxx154 | Expiration date \(applies to all domains and services\)                                                                                 |
| xxx155    | Login ID |
| xxx156 | Password \(use this for domain name account password and domain password; xxx303 is Web hosting account password\)                                                    |
| xxx157    | Subaccount\(s\) \(generic\) |
| xxx158 | Name server name                                                                                                     |
| xxx159    | Name server IP address |
| xxx160 | End user IP address                                                                                                   |
| xxx161    | Server IP address |
| xxx162 | Parking page                                                                                                       |
| xxx163    | Renewal settings |
| xxx164 | Notification amount                                                                                                   |
| xxx165    | Customer preferences |
| xxx166 | Connection timeout                                                                                                    |
| xxx167    | Session timeout |
| xxx168 | Account type                                                                                                       |
| xxx169    | Domain count |
| xxx170 | Email record\(s\)                                                                                                    |
| xxx171    | Host records\(s\) |
| xxx172 | Forward to                                                                                                        |
| xxx173    | URL address \(forwarding—URL address in host records\) |
| xxx174 | Price\(s\)                                                                                                        |
| xxx175    | Queue |
| xxx176 | Club Drop                                                                                                        |
| xxx177    | Reseller key |
| xxx178 | Reseller pricing                                                                                                     |
| xxx179    | Retail pricing |
| xxx180 | Price break                                                                                                       |
| xxx181    | Account balance |
| xxx182 | Product type                                                                                                       |
| xxx183    | Quantity |
| xxx184 | Reference\(s\)                                                                                                      |
| xxx185    | Authorization answer |
| xxx186 | Order number                                                                                                       |
| xxx187    | Reference number |
| xxx188 | Page                                                                                                           |
| xxx189    | Page size |
| xxx190 | URL \(generic, including paging\)                                                                                            |
| xxx191    | Address book |
| xxx192 | Shopping cart                                                                                                      |
| xxx193    | Item ID |
| xxx194 | Host name \(name of a host record, not a Web hosting account\)                                                                              |
| xxx195    | Address |
| xxx196 | Month                                                                                                          |
| xxx197    | Year\(s\) |
| xxx198 | Card type                                                                                                        |
| xxx199    | Cardholder name |
| xxx200 | Cardholder address                                                                                                    |
| xxx201    | Cardholder city |
| xxx202 | Cardholder state/province                                                                                                |
| xxx203    | Cardholder postal code |
| xxx204 | Cardholder country                                                                                                    |
| xxx205    | Credit card number |
| xxx206 | Cardholder phone                                                                                                     |
| xxx207    | Amount |
| xxx208 | Address does not match                                                                                                  |
| xxx209    | Postal code does not match |
| xxx210 | Credit card verification number                                                                                             |
| xxx217    | Date |
| xxx218 | Metatag\(s\)                                                                                                       |
| xxx219    | Watch list |
| xxx220 | Service\(s\)                                                                                                       |
| xxx221    | Transfer order\(s\) |
| xxx222 | Transfer order detail\(s\)                                                                                                |
| xxx223    | Web Site Builder |
| xxx224 | Information                                                                                                       |
| xxx225    | Option ID |
| xxx226 | Token                                                                                                          |
| xxx227    | Count |
| xxx228 | Registrar Lock                                                                                                      |
| xxx229    | Domain name ID |
| xxx230 | Credit card processing                                                                                                  |
| xxx231    | Credit card information |
| xxx232 | Contact \(generic\)                                                                                                   |
| xxx233    | Item count |
| xxx234 | Page view\(s\)                                                                                                      |
| xxx235    | Bundle \(generic—use for .name bundle, and so on\) |
| xxx236 | Item\(s\)                                                                                                        |
| xxx237    | Character limit |
| xxx238 | Word\(s\)                                                                                                        |
| xxx239    | Status |
| xxx240 | Commission balance                                                                                                    |
| xxx241    | Load parameter |
| xxx242 | Order\(s\)                                                                                                        |
| xxx243    | RequestID |
| xxx244 | Record\(s\)                                                                                                       |
| xxx245    | Preference\(s\) |
| xxx246 | Time stamp                                                                                                        |
| xxx247    | Reseller API |
| xxx248 | Reseller                                                                                                         |
| xxx249    | Active reseller |
| xxx250 | Usage                                                                                                          |
| xxx251    | Setting\(s\) |
| xxx252 | Request\(s\)                                                                                                       |
| xxx253    | Parking text |
| xxx254 | User\(s\)                                                                                                        |
| xxx255    | Account\(s\) |
| xxx256 | Option                                                                                                          |
| xxx257    | Order type |
| xxx258 | Authorization                                                                                                      |
| xxx259    | Order ID |
| xxx260 | Detail ID                                                                                                        |
| xxx261    | Action |
| xxx262 | Comment\(s\)                                                                                                       |
| xxx263    | Display flag |
| xxx264 | Object\(s\)                                                                                                       |
| xxx265    | Date |
| xxx266 | Image\(s\)                                                                                                        |
| xxx267    | Preconfigure DNS |
| xxx268 | Submit times                                                                                                       |
| xxx269    | Pak ID of a POP pak associated with a domain name in a domain name account. For POP units associated with a Web hosting account, see xxx314. For package IDs of products other than domain name POP paks, see xxx327. |
| xxx270 | Meta base                                                                                                        |
| xxx271    | POP pak associated with a domain name in a domain name account. For POP units associated with a Web hosting account, see xxx314. |
| xxx272 | Full name                                                                                                        |
| xxx273    | Runtime |
| xxx274 | Product ID                                                                                                        |
| xxx275    | Bid |
| xxx276 | Auction                                                                                                         |
| xxx277    | Site ID |
| xxx278 | Location name                                                                                                      |
| xxx279    | Group name |
| xxx280 | URI                                                                                                           |
| xxx281    | Time zone |
| xxx282 | Start time                                                                                                        |
| xxx283    | End time |
| xxx284 | Alert interval                                                                                                      |
| xxx285    | Name server\(s\) |
| xxx286 | User name of a POP mailbox in a POP pak associated with a domain name. For a Web-hosting-account mailbox, see xxx313.                                                |
| xxx287    | Keyword \(generic\) |
| xxx288 | Type \(generic\)                                                                                                     |
| xxx289    | Extended attribute |
| xxx290 | Contact information \(generic\)                                                                                             |
| xxx291    | Cookie |
| xxx292 | MX Preference                                                                                                      |
| xxx293    | Record type |
| xxx294 | Accounting transaction                                                                                                  |
| xxx295    | Account information |
| xxx296 | TLD information                                                                                                     |
| xxx297    | Site processor |
| xxx298 | Registrar information                                                                                                  |
| xxx299    | Database |
| xxx300 | Hosts count \(number of Web hosting accounts in a domain name account\)                                                                         |
| xxx301    | Web hosting account directory |
| xxx302 | Web hosting account user name                                                                                              |
| xxx303    | Web hosting account password |
| xxx304 | Web hosting host header\(s\)                                                                                               |
| xxx305    | Domain header\(s\) |
| xxx306 | File permissions                                                                                                     |
| xxx307    | Files |
| xxx308 | SQL table\(s\)                                                                                                      |
| xxx309    | Server information |
| xxx310 | Bandwidth                                                                                                        |
| xxx311    | Web hosting account ID |
| xxx312 | Database storage                                                                                                     |
| xxx313    | Individual POP3 mailboxes in a Web hosting account. For POP paks associated with a domain name in a domain name account, see xxx271. |
| xxx314 | POP paks \(units of 10\) in a Web hosting account. For POP paks associated with a domain name in a domain name account, see xxx271.                                          |
| xxx315    | Web storage |
| xxx316 | Overage option                                                                                                      |
| xxx317    | Stored procedure name |
| xxx318 | SQL statement                                                                                                      |
| xxx319    | Table name |
| xxx320 | Index name                                                                                                        |
| xxx321    | Database column |
| xxx322 | Database type                                                                                                      |
| xxx323    | Operating system type |
| xxx324 | Web hosting account ID                                                                                                  |
| xxx325    | Path name |
| xxx326 | Package\(s\) Use for all packages except POP paks associated with a domain name in a domain name account. For domain-name POP paks, see xxx271.                                    |
| xxx327    | Package ID Use for all products sold as multiple-unit “packages” except POP paks associated with a domain name in a domain name account. For the Pak ID of a domain-name POP pak, see xxx269. |
| xxx328 | Package name                                                                                                       |
| xxx329    | Stored procedure text |
| xxx330 | Stored procedure\(s\)                                                                                                  |
| xxx331    | Index property\(s\) |
| xxx350 | Agreement                                                                                                        |
| xxx351    | Secret type |
| xxx352 | Display name                                                                                                       |
| xxx353    | Location count |
| xxx354 | Notification count                                                                                                    |
| xxx355    | ExpiredCount—number of expired domain names currently in this account. |
| xxx356 | DomainCount\_Others                                                                                                   |
| xxx357    | ExpiredCount\_Others |
| xxx358 | Language code                                                                                                      |
| xxx359    | Day\(s\) |
| xxx360 | Order status                                                                                                       |
| xxx361    | Email key\(s\) |
| xxx362 | Registrar                                                                                                        |
| xxx363    | Color |
| xxx364 | Fraud record\(s\)                                                                                                    |
| xxx365    | Fraud type |
| xxx366 | Fraud value                                                                                                       |
| xxx367    | Fraud level |
| xxx368 | Enable instruction\(s\)                                                                                                 |
| xxx369    | PayPal contract |
| xxx370 | Category ID                                                                                                       |
| xxx371    | Subcategory ID |
| xxx372 | Category name                                                                                                      |
| xxx373    | Subcategory name |
| xxx374 | Category keyword type code                                                                                                |
| xxx375    | Filter mask |
| xxx376 | Category keyword reference ID                                                                                              |
| xxx377    | Category keyword relevance level |
| xxx378 | Access control violation -- Stored procedure failure                                                                                  |
| xxx379    | Access control violation -- Access control lock |
| xxx380 | Access control violation -- Domain not under account                                                                                  |
| xxx381    | Access control violation -- Escrow hold |
| xxx382 | Access control violation -- Invalid PartyID                                                                                      |
| xxx383    | Access control violation -- Invalid access type |
| xxx384 | Invalid promotion code                                                                                                  |
| xxx385    | Adult |
| xxx386 | Hard core                                                                                                        |
| xxx387    | China |
| xxx388 | Web Site Creator                                                                                                     |
| xxx389    | Source |
| xxx390 | Reason                                                                                                          |
| xxx391    | bundleID |
| xxx392 | Device                                                                                                          |
| xxx393    | Folder |
| xxx394 | Hosted Microsoft Exchange                                                                                                |
| xxx396    | ROID |
| xxx397 | SSL Certificate                                                                                                     |
| xxx398    | Channel |
| xxx399 | Social Networking                                                                                                    |
| xxx400    | Domain Spinner |
| xxx401 | Membership                                                                                                        |
| xxx402    | Business Listing |
| xxx403 | Mobilizer                                                                                                        |
| xxx404    | Rich Content |
| xxx405 | RatePoint                                                                                                        |
| xxx406    | Tag |
| xxx999 | Unknown                                                                                                         |

Email error codes
-----------------

> ### Password Validation
>
>
>
> Password Validation can only occur when a plaintext password is sent as part of the API request:
>
>
>
> A user changing their own password is subject to an additional password strength test. Increasing the length of the password and including a mix of symbols, upper and lower case letters, and numbers will increase the calculated strength of a given password.
>
>
>
> All password changes are subject to the following check:
>
>
>
> - Contains check: any password supplied during a password update or account creation will be checked to make sure it does not contain the user's username or domain name, and if it does, the password will be rejected. This test is not case-sensitive.

> ### When creating a new user, if a service is not specified in the request, the value of the service attribute is taken from the domain setting. If the domain does not have a value for that service, it is enabled. If a service is suspended in the domain, it can only be set in the request by an admin of the same or higher level than the admin that set the service to suspended in the domain \(for example, a domain admin cannot change a service that was suspended by a company admin\).

| Error code | Error text | Definition                                                                                                 |
| ---------- | ------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 0     | Server error. | Returned when something goes wrong and it is not the client's fault.                                                                   |
| 1     | Invalid credentials supplied in request. | Returned when the user name or password is not valid.                                                                           |
| 2     | The requested object does not exist. | Returned when the object \(user, domain, brand, etc\) being queried, changed or deleted does not exist.                                                  |
| 3     | This object is an alias. | Returned when the name of an object to be changed in a request is an alias or an object. Objects can only be changed via their real name.                                |
| 4     | Requestor lacks permission to change one or more of the requested attributes. | Returned when the request specifies a change to an attribute that the user does not have permission to change.                                              |
| 5     | Request badly formatted \(missing required field, or field is not the correct data type\). | Returned when the request is not properly formatted.                                                                           |
| 6     | One or more attributes badly formatted. | Returned when the request to change an object contains an unrecognized attribute or an attribute value that is not well formatted. For example, the value of the key filter might not be a valid number. |
| 7     | An object with this name already exists. | Returned when the request to create an object names an object that already exists.                                                            |
| 8     | An object with this name already exists. | Returned when the request to create an object names an object that already exists.                                                            |
| 9     | Requestor does not own this object or lacks permission to perform this action. | Returned when the user lacks permission to perform the requested action.                                                                 |
| 10     | The requested object is not empty. | Returned when the request asks to delete an object, but the object is not empty \(for example to delete a domain but the domain contains users\).                             |
| 11     | Company does not exist. | Returned with the company containing an object to be changed \(domain, brand\) does not exist.                                                      |
| 12     | Role does not exist | Returned when attempting to add a non-existent role to a user.                                                                      |
| 13     | User does not exist | Returned when attempting to do something with a user \(post a bulletin to a test user\) and that user does not exist.                                           |
| 14     | Brand in use | Returned when attempting to delete a brand, and that brand is assigned to a user, domain or company.                                                   |
| 15     | Domain users full | Returned when attempting to create a user, and the user's domain is at its user limit.                                                          |
| 16     | Domain aliases full | Returned when attempting to create an alias to a user and the domain is at its alias limit.                                                        |
| 17     | Not in | Returned when an operation requires a user to be an member of something and the user is not.                                                       |
| 18     | Workgroup is default | Returned when attempting to delete a workgroup that is the domain's default workgroup.                                                          |
| 19     | Migration Job Exists | Returned when attempting to create a new migration job, and a migration job with the given ID already exists.                                               |
| 20     | Try again later | A transient system error occurred. Not the caller's fault. Method should be retried.                                                          |
| 21     | Requestor not permitted from this IP address | IP restrictions have been set for this admin user and the source IP does not match them.                                                         |
| 22     | Requested object temporarily unavailable, try again later | Temporary error, please retry in a few minutes.                                                                              |
| 23     | Object already exists | If object already exists and create\_only attribute was passed in the request.                                                              |
| 24     | Admins exist that control this object |                                                                                                      |
| 25     | Failed to obtain provisioning lock | Temporary error, please retry in a few minutes.                                                                              |
| 26     | The requested folder does not exist |                                                                                                      |
| 27     | The domain is missing the filtermx setting |                                                                                                      |
| 28     | An internal error occurred | If error persists, please contact Enom support                                                                               |
| 29     | An internal error occurred - failed to connect | If error persists, please contact Enom support                                                                               |
| 30     | An internal error occurred - failed to send | If error persists, please contact Enom support                                                                               |
| 31     | User was not type filter |                                                                                                      |
| 32     | Admin is disabled |                                                                                                      |
| 33     | 2FA is required | Returned when second stage Two-Factor Authentication credentials are required                                                               |
| 34     | password must be at least 10 characters | Returned when using reseller\_id credentials during 2FA signup if password is not deemed secure enough                                                   |
| 35     | 2fa code validation failure | Returned when second stage 2FA credentials are incorrect                                                                          |
| 36     | password cannot be one of the previous 8 used | Returned when using reseller\_id credentials during 2FA signup and new password was previously used                                                    |
| 37     | Alias domain does not exist |                                                                                                      |
| 38     | maximum number of forward recipients reached | Returned when trying to set more forward recipients than allowed on an account.                                                              |
| 39     | expired |                                                                                                      |
| 40     | Password not strong enough | Returned when a user's password fails the password strength check \(see change\_user\)                                                           |
| 42     | 2FA is required, not enabled on this account | Returned during authentication when 2FA is forced for user, user has not enabled it, and user is out of grace logins.                                           |
| 43     | 2FA challenge failed | Returned when second stage 2FA credentials are incorrect                                                                          |
| 44     | 2FA Auth required to change 2fa\_required\_oma | Attribute can only be changed when user authenticated with 2FA                                                                       |
| 45     | Brand settings prevent this action |                                                                                                      |
| 46     | Max number of App Passwords has been reached | Target user has reached maximum \(4\) number of App Passwords.                                                                      |
| 47     | Target is a domain alias | Cannot rename a user into a domain alias; can only rename the user into the true domain name.                                                       |
| 48     | Password rejected - contains username or other account info | Returned when a user's password fails the password contains check \(see change\_user\)                                                           |

 