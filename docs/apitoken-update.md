ApiToken\Update
===============

Update the name of an API token.

Usage
-----

Use this command when you want to update the name of an API token.

Availability
------------

All resellers have access to this command.

Constraints
-----------
- The name must be unique.
- The token id is required.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=ApiTokens_update&uid={YourAccountID}&pw={YourAccountPassword}&tokenid={tokenid}&responsetype=xml
```
| Input Parameter   | Type | Description         |
| -------------------- | ------- | ---------------------------- |
| command       | string | **ApiToken\_Update**     |
| uid         | string | Your account username    |
| pw          | string | Your API token        |
| tokenid \(required\) | integer | The API token ID to update  |
| name \(required\)  | integer | The API token name to update |
| responsetype     | string | **XML**           |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined above.

The following params are nested in:

```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
  <Success>True</Success>
  <ErrCount>0</ErrCount>
</interface-response>
```
| Output Parameter | Type | Description  |
| ---------------- | ---- | ------------- |
| Success     | bool | True or False |