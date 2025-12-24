ApiToken\Details
================

Retrieve one or all API tokens.

Usage
-----

Use this command when you want to retrieve one or all API tokens associated with a user.

Availability
------------

All resellers have access to this command.

Constraints
-----------

You will need to supply your account password as the **pw** parameter to execute this command.

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=ApiToken_Details&uid={YourAccountID}&pw={YourAccountPassword}&responsetype=xml
```
| Input Parameter  | Type | Description                                       |
| ------------------ | ------ | ---------------------------------------------------------------------------------------- |
| command      | string | **ApiToken\_Details**                                  |
| uid        | string | Your account username                                  |
| pw         | string | Your account password                                  |
| token \(optional\) | string | The API token you want to retrieve. If not supplied, all API tokens will be retrieved. |
| responsetype    | string | **XML**                                         |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined above.

The following params are nested in:

```
<Tokens>
    <Token>
      <APIToken>...</APIToken>
      <APITokenName>...</APITokenName>
      <APITokenId>...</APITokenId>
    </Token>
  </Tokens>
```
| Output Parameter | Type | Description                   |
| ---------------- | ------ | ----------------------------------------------- |
| APIToken     | string | The API token                  |
| APITokenName   | string | The API token name               |
| APITokenID    | int | The API token ID \(used for internal tracking\) |