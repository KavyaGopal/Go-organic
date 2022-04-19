import { Component, OnInit } from '@angular/core';
import { render } from 'creditcardpayments/creditCardPayments' 
@Component({
  selector: 'app-paypalpayment',
  templateUrl: './paypalpayment.component.html',
  styleUrls: ['./paypalpayment.component.css']
})
export class PaypalpaymentComponent implements OnInit {

  constructor() {
    render(
      {
        id:"#myCheckoutPayPal",
        currency:"USD",
        value:"100.00",
        onApprove:(details)=> {
          alert("Transaction Successful. You may close the browser now.");

        }
      }
    );
   }

  ngOnInit(): void {
  }

}
