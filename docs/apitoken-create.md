ApiToken\Create
===============

Generate an API Token.

Usage
-----

Generate a unique API token to be used for API authentication. The API token will replace your account password in your API calls.

Availability
------------

All resellers have access to this command.

Constraints
-----------

You will need to supply your account password as the **pw** parameter to execute this command.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=ApiToken_Create&uid={YourAccountID}&pw={YourAccountPassword}&name={UniqueName}&responsetype=xml
```
| Input Parameter | Type | Description                   |
| --------------- | ------ | ------------------------------------------------ |
| command     | string | **ApiToken\_Create**               |
| uid       | string | Your account username              |
| pw       | string | Your account password              |
| name      | string | A **unique name** used to identify the API token |
| responsetype  | string | **XML**                     |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined above.

| Output Parameter | Type | Description    |
| ---------------- | ------ | ------------------ |
| APIToken     | string | The API Token   |
| APITokenName   | string | The API Token Name |
| APITokenID    | int | The API Token ID  |