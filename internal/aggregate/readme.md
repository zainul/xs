# Aggregate

Aggregate is actually the package that contain entity that aggregate or joining with other entity,
for example :

You have 2 entity history and user, in history you want to join with the user.
in simple case you can join with history and user , but for the next level , user will be use different database with the history
for example :

**user in mongo**
**history in postgres**

in gateway level you cannot join the both entity. Actually the aggregation can appear in repository level it's quite same, 
but why need package aggregate ? is because to be loosely couple between entity and the gateway (we can added any custom aggregate structure base on business need)