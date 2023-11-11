import { Component } from '@angular/core';

@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css']
})
export class UserComponent {
  data = [
    { date:" 11/2/2023", time: '11 pm', doctor: "ahmed" },
    { date: "11/4/2023", time: '10 am', doctor: "sakr"},
    { date: "11/7/2023", time: '6 pm', doctor: "wael" }
  ];
  

}
