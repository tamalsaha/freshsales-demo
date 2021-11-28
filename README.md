# freshsales-demo

```bash
curl -H "Authorization: Token token=$FRESHSALES_TOKEN" -H "Content-Type: application/json" -X GET "https://$FRESHSALES_DOMAIN.myfreshworks.com/crm/sales/api/settings/contacts/fields"

curl -H "Authorization: Token token=$FRESHSALES_TOKEN" -H "Content-Type: application/json" -X GET "https://$FRESHSALES_DOMAIN.myfreshworks.com/crm/sales/api/contacts/filters"

# All Contacts view_id=17002064143
curl -H "Authorization: Token token=$FRESHSALES_TOKEN" -H "Content-Type: application/json" -X GET "https://$FRESHSALES_DOMAIN.myfreshworks.com/crm/sales/api/contacts/view/17002064143"
```
