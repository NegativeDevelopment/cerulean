import 'package:flutter/material.dart';

class HomePage extends StatefulWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          title: Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              Row(
                children: [
                  IconButton(
                    icon: const Icon(Icons.menu),
                    onPressed: () {},
                  ),
                  Text('Cerulean', style: Theme.of(context).textTheme.headline6)
                ],
              ),
              Text("Username", style: Theme.of(context).textTheme.labelLarge)
            ],
          ),
        ),
        body: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 12.0, vertical: 8.0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              // Group card
              Card(
                  child: Padding(
                padding: const EdgeInsets.symmetric(
                    horizontal: 16.0, vertical: 12.0),
                child: Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    Column(
                      crossAxisAlignment: CrossAxisAlignment.start,
                      children: [
                        Text("Group 1",
                            style: Theme.of(context).textTheme.headline6),
                        Text("102 members",
                            style: Theme.of(context).textTheme.bodyText2),
                        Text("50 Transactions",
                            style: Theme.of(context).textTheme.bodyText2),
                      ],
                    ),
                    ClipRRect(
                      borderRadius:
                          const BorderRadius.all(Radius.circular(8.0)),
                      child: Image.asset(
                        "assets/group-placeholder.jpg",
                        width: 60,
                        height: 60,
                        fit: BoxFit.cover,
                      ),
                    )
                  ],
                ),
              )),

              //Offene Payments
              SizedBox(
                height: 200,
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Padding(
                      padding: const EdgeInsets.symmetric(vertical: 16.0),
                      child: Text("Pending Payments",
                          style: Theme.of(context).textTheme.headline6),
                    ),
                    Expanded(
                      child: ListView.builder(
                        shrinkWrap: false,
                        scrollDirection: Axis.horizontal,
                        itemCount: 20,
                        itemBuilder: (context, index) {
                          // return Container(
                          //   width: 100,
                          //   height: 100,
                          //   margin: EdgeInsets.all(8.0),
                          //   color: Colors.red,
                          // );
                          return Container(
                            margin: const EdgeInsets.symmetric(horizontal: 8.0),
                            width: 60,
                            child: Column(children: const [
                              Icon(Icons.account_circle_outlined, size: 60),
                              Text(
                                "Username 123 123 123",
                                overflow: TextOverflow.ellipsis,
                                maxLines: 1,
                              ),
                              Text(
                                "+50 €",
                                maxLines: 1,
                                overflow: TextOverflow.fade,
                              )
                            ]),
                          );
                        },
                      ),
                    ),
                  ],
                ),
              ),

              // Transactions
              Expanded(
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Padding(
                      padding: const EdgeInsets.symmetric(vertical: 16.0),
                      child: Text("Transactions",
                          style: Theme.of(context).textTheme.headline6),
                    ),
                    Expanded(
                        child: ListView.separated(
                            // shrinkWrap: true,
                            itemCount: 100,
                            separatorBuilder: (context, index) {
                              return const Divider();
                            },
                            itemBuilder: (context, index) {
                              return ListTile(
                                title: Text("Transaction $index"),
                                subtitle: Text("Subtitel $index"),
                                leading: Icon(
                                  Icons.ev_station,
                                  color: Theme.of(context).colorScheme.primary,
                                  size: 40,
                                ),
                                trailing: const Text("2Dollar 50"),
                              );
                            })),
                  ],
                ),
              )
            ],
          ),
        ));
  }
}
