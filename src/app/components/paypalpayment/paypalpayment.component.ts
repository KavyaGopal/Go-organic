import { Component, OnInit } from '@angular/core';
import { render } from 'creditcardpayments/creditCardPayments' 
import {Router} from '@angular/router'; // import router from angular router

@Component({
  selector: 'app-paypalpayment',
  templateUrl: './paypalpayment.component.html',
  styleUrls: ['./paypalpayment.component.css']
})
export class PaypalpaymentComponent implements OnInit {

  constructor(private route:Router) {
    render(
      {
        id:"#myCheckoutPayPal",
        currency:"USD",
        value:"100.00",
        onApprove:(details)=> {
          alert("Transaction Successful. You may close the browser now.");
          this.route.navigate(['/payment']);

        }
      }
    );
   }

  ngOnInit(): void {
  }

}
