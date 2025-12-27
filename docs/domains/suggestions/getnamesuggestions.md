GetNameSuggestions
==================

Generate variations of the domain name you specify.

Usage
-----

Use this command to generate variations of the domain name you specify. This command will return a list of domains across a variety of SLD/TLD combinations.

Availability
------------

All resellers have access to this command.

Constraints
-----------

None

Input Parameters
----------------

```
https://resellertest.enom.com/interface.asp?command=GetNameSuggestions&uid=YourAccountID&pw=YourApiToken&{param1}={value1}&responsetype=xml
```
| Input Parameter | Type | Description                                                                                                                                                       |
| --------------- | ------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| command     | string | GetNameSuggestions                                                                                                                                                    |
| uid       | string | Your Account ID                                                                                                                                                     |
| pw       | string | Your API Token                                                                                                                                                      |
| SearchTerm   | string | Required - The term to use to generate suggestions, most commonly use is SLD                                                                                                                      |
| TldList     | string | A coma separated list of TLDs to generate results on. Only TLDs included in this list will be included in the results. When used with AllGA it will also include additional TLDs you specify, such as some in pre-register status if desired.                                    |
| OnlyTldList   | string | A coma separated list of TLDs to generate results on. Only TLDs included in this list will be included in the results, even if AllGA is true.                                                                                     |
| ExcludeTldList | string | TLDs A coma separated list of TLDs to exclude from the results.                                                                                                                             |
| MaxResult    | int | Maximum number of suggested names to return in addition to your input. The number actually returned may be lower based on search term. Permitted values are numbers 1 through 100                                                                   |
| SpinType    | int | Determine the type of results you will get; 0 - Both TLD and Spun recommendations - combination of 1 & 2 1 - Only TLD recommendations \(exact SLD match\) - no suggestions 2 - Only Spun recommendations - only suggestions 3 - Will allow suggestions such as domaintld to be treated as domain.tld suggestions |
| Adult      | boolean | Allow domain names that may be considered "adult" in nature. Permitted values are True or False.                                                                                                            |
| Premium     | boolean | Allow suggestions that are considered to be "premium" to be returned in the result set. Permitted values are True or False                                                                                               |
| AllGA      | boolean | Controls if only TLDs in GA \(General Availability\) are returned, or if it will return ALL TLDs, including those in pre-registration. Permitted values are True or False.                                                                      |

Returned Parameters and Values
------------------------------

The default response format is plain text. To receive the response in HTML or XML format, send ResponseType=HTML or ResponseType=XML in your request.

Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined above.

| Output Parameter    | Type | Description                                                                   |
| ---------------------- | ------- | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| DomainSuggesstionCount | int | Number of Suggestions returned in this response                                                 |
| Score         | double | Relevancy of the result to the search term                                                    |
| Tld          | string | TLD for this set of results                                                           |
| Sld          | string | SLD for this set of results                                                           |
| Idn          | boolean | Is this domain an IDN. True indicates yes; False indicates no.                                         |
| Premium        | boolean | Is this domain considered "Premium"? True indicates yes; False indicates no. Unknown indicates that this is a TLD which we cant know for sure. |
| In\_GA         | boolean | Is this TLD in GA. True indicates yes; False indicates no.                                           |
| NativeSld       | string | If the item is an IDN, the Native language SLD.                                                 |
| NativeTld       | string | If the item is an IDN, the Native language TLD.                                                 |