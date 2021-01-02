#D2H Problem Statement

- Direct to home (D2H) operator SatTV wants you to design a software system for
its customers.
- User should have an initial balance of 100 Rs. in the account.
- User can view the current subscription and balance
- User can recharge account with any positive amount.
- SatTV offers multiple channel packages (called base packs henceforth). Base
packs available to purchase are Gold and Silver, which comes with certain
channels.
- User can view the details of available packs and the channels.
- Whenever a user subscribes to the base pack, the base pack price will be
deducted from the account balance.
- User should receive a 10% discount on all the base packs if the duration is at
least 3 months.
- User can add the individual channel into the current subscription on which
amount will be deducted from the account balance.
- Whenever the back pack is subscribed or if any channel is added in the base
pack, notifications should be sent through email and/or SMS if the user has
updated email and/or phone number.

# Sample Input/Output for Console Application

Welcome to SatTV
1. View account balance and current subscription
2. Recharge Account
3. View available packs and channels
4. Subscribe to base packs
5. Add channels to an existing subscription
6. Exit

Enter the option: 1
Account balance is 100 Rs.
Current subscription : Gold (If the user has not yet subscribed show “ Not
subscribed” )

Enter the option: 2
Enter the amount to recharge: 500
Recharge completed successfully. Current balance is 600

Enter the option: 3
Available packs for subscription
Silver pack: Zee, Sony, Star Plus: 50 Rs.
Gold Pack: Zee, Sony, Star Plus, Discovery, NatGeo: 100 Rs.
Available channels for subscription
Zee: 10 Rs, Sony: 15 Rs, Star Plus: 20 Rs, Discovery: 10 Rs, NatGeo: 20 Rs.

Enter the option: 4
Enter the Pack you wish to subscribe: (Silver: ‘S’, Gold: ‘G’): G
Enter the months: 3
You have successfully subscribed the following pack - Gold
Monthly price: 100 Rs.
No of months: 3
Subscription Amount: 300 Rs.
Discount applied: 30 Rs.
Final Price after discount: 270 Rs
Email notification sent successfully
SMS notification sent successfully

Enter the option: 5
Enter channel names to add (separated by commas): Discovery, Nat Geo
Channels added successfully.
Account balance: 300 Rs.

Enter the option: 6
Exit.
