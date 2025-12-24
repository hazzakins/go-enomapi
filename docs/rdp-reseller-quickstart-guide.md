RDP Reseller Quickstart Guide
=============================

Work in Progress

**Note - these changes are currently available only on Reseller Test for .org and .asia TLDs**

This guide is to help you update your system for the new ICANN RDP rules going into effect this year. Although we recommend that resellers collect the minimum data set, we will not reject orders if a reseller collects and provides us with additional data beyond this minimum data set. However, we will not retain or store any additional data beyond the minimum data set and specific additional data required by the registry for that TLD.

GetTldDetails(../docs/get-tld-details.html)
-------------------------------------------

Use this command to determine if RDP is enabled and if enabled what rules apply:

[https://resellertest.enom.com/interface.asp?command=gettlddetails&uid=\{YourAccountID\}&pw=\{YourPassword\}&responsetype=xml&tld=org](https://resellertest.enom.com/interface.asp?command=gettlddetails&uid=%7BYourAccountID%7D&pw=%7BYourPassword%7D&responsetype=xml&tld=org)

```
<tlds>
  <tld>
  <TLDID>2</TLDID>
  <TLD>org</TLD>
  <AbleToLock>True</AbleToLock>
  <AllowWPPS>True</AllowWPPS>
  <AutoRenewOnly>False</AutoRenewOnly>
  <ExtAttributes>False</ExtAttributes>
  <AllowWBL>True</AllowWBL>
  <TransRequiresNewContact>False</TransRequiresNewContact>
  <EncodingType>PUNY</EncodingType>
  <LangCode/>
  <NativeDisplay/>
  <ClearPhoneNumber>False</ClearPhoneNumber>
  <ClearFax>False</ClearFax>
  <ValidateDNSHosting>True</ValidateDNSHosting>
  <HasPremiumNames>False</HasPremiumNames>
  <SupportsDnsSec>True</SupportsDnsSec>
  <AllowWhoisPublicity>True</AllowWhoisPublicity>
  <RDPEnabled>True</RDPEnabled>
```
The RDPEnabled flag will tell you if RDP is enabled for that TLD.

You can then look at the RDPRules node to see the rules \(node will be when RDPEnabled is false\):

```
	ContactType: What contact type the rule affects.
	ContactField: What contact field name to which the rule applies.
	CollectionType:
			required - you must send this field or will fail at the registry if applicable.
			optional - if you send this information it will be in the contact.
	TransmitToRegistry:
			True - we send that contact and field name to the registry
			False - we do not send that contact and field name to the registry.
```
Example - .org call tells you that we only create registrant contact \(because no rules for admin, tech or billing are returned\), what fields we use to create the contact and that none of them are sent to the registry \(.org is going thin when they launch RDP and their OTE reflects this change\).

```
<RDPRules>
   <RDPRule>
     <ContactType>registrant</ContactType>
     <ContactField>First Name</ContactField>
     <CollectionType>required</CollectionType>
     <TransmitToRegistry>False</TransmitToRegistry>
   </RDPRule>
   <RDPRule>
     <ContactType>registrant</ContactType>
     <ContactField>Last Name</ContactField>
     <CollectionType>required</CollectionType>
     <TransmitToRegistry>False</TransmitToRegistry>
   </RDPRule>
   <RDPRule>
     <ContactType>registrant</ContactType>
     <ContactField>Organization</ContactField>
     <CollectionType>optional</CollectionType>
     <TransmitToRegistry>False</TransmitToRegistry>
   </RDPRule>
   <RDPRule>
     <ContactType>registrant</ContactType>
     <ContactField>Address</ContactField>
     <CollectionType>required</CollectionType>
     <TransmitToRegistry>False</TransmitToRegistry>
   </RDPRule>
   <RDPRule>
     <ContactType>registrant</ContactType>
     <ContactField>Address 2</ContactField>
     <CollectionType>optional</CollectionType>
     <TransmitToRegistry>False</TransmitToRegistry>
   </RDPRule>
   <RDPRule>
     <ContactType>registrant</ContactType>
     <ContactField>City</ContactField>
     <CollectionType>required</CollectionType>
     <TransmitToRegistry>False</TransmitToRegistry>
   </RDPRule>
   <RDPRule>
     <ContactType>registrant</ContactType>
     <ContactField>State or Province</ContactField>
     <CollectionType>optional</CollectionType>
     <TransmitToRegistry>False</TransmitToRegistry>
   </RDPRule>
   <RDPRule>
     <ContactType>registrant</ContactType>
     <ContactField>Postal Code</ContactField>
     <CollectionType>optional</CollectionType>
     <TransmitToRegistry>False</TransmitToRegistry>
   </RDPRule>
   <RDPRule>
     <ContactType>registrant</ContactType>
     <ContactField>Country Code</ContactField>
     <CollectionType>required</CollectionType>
     <TransmitToRegistry>False</TransmitToRegistry>
   </RDPRule>
   <RDPRule>
     <ContactType>registrant</ContactType>
     <ContactField>Phone Number</ContactField>
     <CollectionType>required</CollectionType>
     <TransmitToRegistry>False</TransmitToRegistry>
   </RDPRule>
   <RDPRule>
     <ContactType>registrant</ContactType>
     <ContactField>Email Address</ContactField>
     <CollectionType>required</CollectionType>
     <TransmitToRegistry>False</TransmitToRegistry>
   </RDPRule>
   <RDPRule>
     <ContactType>registrant</ContactType>
     <ContactField>Fax Number</ContactField>
     <CollectionType>optional</CollectionType>
     <TransmitToRegistry>False</TransmitToRegistry>
   </RDPRule>
 </RDPRules>
```
Purchase(../docs/purchase.html)
-------------------------------

The command works the same as before except for RDP tlds the rules are applied on contact creation. You can use the rules to limit what you are sending but we will just ignore anything data sent that is not required to create RDP contact.

Example - .org Purchase returns empty nodes for Admin, Tech and Billing because no contacts were created for those types.

```
<Contacts>
  <Registrant>
    <OrganizationName>Business Ltd</OrganizationName>
    <FirstName>First</FirstName>
    <LastName>Name</LastName>
    <Address1>1232 Testing Ave.</Address1>
    <Address2/>
    <City>Trento</City>
    <StateProvince>AL</StateProvince>
    <PostalCode>38100</PostalCode>
    <Country>US</Country>
    <Phone>+39.046455555</Phone>
    <Fax/>
    <EmailAddress>[email protected]</EmailAddress>
    <ConsentStatus>FORCED_ALL_CONTRACTUAL</ConsentStatus>
  </Registrant>
  <Admin>
    <OrganizationName/>
    <JobTitle/>
    <FirstName/>
    <LastName/>
    <Address1/>
    <Address2/>
    <City/>
    <StateProvince/>
    <PostalCode/>
    <Country/>
    <Phone/>
    <PhoneExt/>
    <Fax/>
    <EmailAddress/>
  </Admin>
  <Tech>
    <OrganizationName/>
    <JobTitle/>
    <FirstName/>
    <LastName/>
    <Address1/>
    <Address2/>
    <City/>
    <StateProvince/>
    <PostalCode/>
    <Country/>
    <Phone/>
    <PhoneExt/>
    <Fax/>
    <EmailAddress/>
  </Tech>
  <Billing>
    <OrganizationName/>
    <JobTitle/>
    <FirstName/>
    <LastName/>
    <Address1/>
    <Address2/>
    <City/>
    <StateProvince/>
    <PostalCode/>
    <Country/>
    <Phone/>
    <PhoneExt/>
    <Fax/>
    <EmailAddress/>
  </Billing>
</Contacts>
```
InsertNewOrder(../docs/InsertNewOrder.html)
-------------------------------------------

The command works the same as before except for RDP tlds the rules are applied on contact creation when the cart order is processed.

Contacts(../docs/contacts.html)
-------------------------------

GetContacts(../docs/getcontacts.html)
-------------------------------------

The command works the same as before except RPD rules are applied so nodes that do not have any information will come back as empty

Example - .org GetContacts returns empty nodes for Admin, Tech and Billing because no contacts were created for those types.

```html
<Registrant>
  <RegistrantPartyID/>
  <RegistrantOrganizationName>
  		<![CDATA[ fake pty ltd ]]>
  </RegistrantOrganizationName>
  <RegistrantJobTitle>
  		<![CDATA[ leader ]]>
  </RegistrantJobTitle>
  <RegistrantFirstName>
  		<![CDATA[ dumb ]]>
  </RegistrantFirstName>
  <RegistrantLastName>
  		<![CDATA[ test ]]>
  </RegistrantLastName>
  <RegistrantAddress1>
  		<![CDATA[ 1232 Testing Ave. ]]>
  </RegistrantAddress1>
  <RegistrantAddress2/>
  <RegistrantCity>
  		<![CDATA[ Trento ]]>
  </RegistrantCity>
  <RegistrantCountry>
  		<![CDATA[ US ]]>
  </RegistrantCountry>
  <RegistrantStateProvince>
  		<![CDATA[ AL ]]>
  </RegistrantStateProvince>
  <RegistrantStateProvinceChoice/>
  <RegistrantPostalCode>
  		<![CDATA[ 38100 ]]>
  </RegistrantPostalCode>
  <RegistrantEmailAddress>
  		<![CDATA[ [email protected] ]]>
  </RegistrantEmailAddress>
  <RegistrantPhone>
  		<![CDATA[ +39.046455555 ]]>
  </RegistrantPhone>
  <RegistrantPhoneExt/>
  <RegistrantFax/>
  <RegistrantConsentStatus>
  		<![CDATA[ FORCED_ALL_CONTRACTUAL ]]>
  </RegistrantConsentStatus>
  <RegistrantConsentStatusDisplay>
  		<![CDATA[ N/A - All Contractual Data ]]>
  </RegistrantConsentStatusDisplay>
  <RegistrantContactMatch>
  		<![CDATA[ UseCustom ]]>
  </RegistrantContactMatch>
</Registrant>
<Admin>
  <AdminPartyID/>
  <AdminOrganizationName/>
  <AdminJobTitle/>
  <AdminFirstName/>
  <AdminLastName/>
  <AdminAddress1/>
  <AdminAddress2/>
  <AdminCity/>
  <AdminCountry/>
  <AdminStateProvince/>
  <AdminStateProvinceChoice/>
  <AdminPostalCode/>
  <AdminEmailAddress/>
  <AdminPhone/>
  <AdminPhoneExt/>
  <AdminFax/>
  <AdminContactMatch>
		  <![CDATA[ NotCollected ]]>
  </AdminContactMatch>
</Admin>
<Tech>
  <TechPartyID/>
  <TechOrganizationName/>
  <TechJobTitle/>
  <TechFirstName/>
  <TechLastName/>
  <TechAddress1/>
  <TechAddress2/>
  <TechCity/>
  <TechCountry/>
  <TechStateProvince/>
  <TechStateProvinceChoice/>
  <TechPostalCode/>
  <TechEmailAddress/>
  <TechPhone/>
  <TechPhoneExt/>
  <TechFax/>
  <TechContactMatch>
		  <![CDATA[ NotCollected ]]>
  </TechContactMatch>
</Tech>
<Billing>
  <BillingPartyID/>
  <BillingOrganizationName/>
  <BillingJobTitle/>
  <BillingFirstName/>
  <BillingLastName/>
  <BillingAddress1/>
  <BillingAddress2/>
  <BillingCity/>
  <BillingCountry/>
  <BillingStateProvince/>
  <BillingStateProvinceChoice/>
  <BillingPostalCode/>
  <BillingEmailAddress/>
  <BillingPhone/>
  <BillingPhoneExt/>
  <BillingFax/>
  <BillingContactMatch/>
</Billing>
  <AuxBilling>
  <AuxBillingPartyID/>
  <AuxBillingOrganizationName/>
  <AuxBillingJobTitle/>
  <AuxBillingFirstName/>
  <AuxBillingLastName/>
  <AuxBillingAddress1/>
  <AuxBillingAddress2/>
  <AuxBillingCity/>
  <AuxBillingCountry/>
  <AuxBillingStateProvince/>
  <AuxBillingStateProvinceChoice/>
  <AuxBillingPostalCode/>
  <AuxBillingEmailAddress/>
  <AuxBillingPhone/>
  <AuxBillingPhoneExt/>
  <AuxBillingFax/>
  <auxID/>
  <regID/>
  <AuxBillingContactMatch>
  		<![CDATA[ NotCollected ]]>
  </AuxBillingContactMatch>
</AuxBilling>
```