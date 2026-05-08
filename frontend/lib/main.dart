import 'package:flutter/material.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    var lessons = [
          {"name": "Math", "time": "12:00"},
          {"name": "Chemistry", "time": "15:00"},
          {"name": "Physics", "time": "10:00"},
        ];
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(
          title: Text("Моё приложение"),
        ),
        

        body: ListView(
          children: lessons.map((lesson)
          {
            return ListTile(
              title: Text(lesson["name"]!),
              subtitle: Text(lesson["time"]!),
            );
          }).toList(),
        ),
      
      ),
    );
  }
}