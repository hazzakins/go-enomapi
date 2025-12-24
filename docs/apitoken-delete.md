ApiToken\Delete
===============

Delete \(revoke\) an API Token.

Usage
-----

Use this command when you want to delete \(revoke\) an API token associated with your account.

Availability
------------

All resellers have access to this command.

Constraints
-----------

You will need to supply your account password as the **pw** parameter to execute this command.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=ApiToken_Delete&uid={YourAccountID}&pw={YourAccountPassword}&token={ApiTokenToDelete}&responsetype=xml
```
| Input Parameter | Type | Description                  |
| --------------- | ------ | --------------------------------------------- |
| command     | string | **ApiToken\_Delete**             |
| uid       | string | Your account username             |
| pw       | string | Your account password             |
| token      | string | The API token you want to delete \(revoke\). |
| responsetype  | string | **XML**                    |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined above.

The following params are nested in:

```
<interface-response>
  <Success>True</Success>
  <ErrCount>0</ErrCount>
</interface-response>
```
| Output Parameter | Type | Description  |
| ---------------- | ---- | ------------- |
| Success     | bool | True or False |