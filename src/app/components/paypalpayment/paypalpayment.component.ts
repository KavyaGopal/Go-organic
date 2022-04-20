import { Component, OnInit } from '@angular/core';
import { render } from 'creditcardpayments/creditCardPayments' 
import { GetApiService } from '../../get-api.service';
@Component({
  selector: 'app-paypalpayment',
  templateUrl: './paypalpayment.component.html',
  styleUrls: ['./paypalpayment.component.css']
})
export class PaypalpaymentComponent implements OnInit {

  constructor(private api:GetApiService) {
    render(
      {
        id:"#myCheckoutPayPal",
        currency:"USD",
        value:"100.00",
        onApprove:(details)=> {
          alert("Transaction Successful. You may close the browser now.");
          this.placeOrder();
        }
      }
    );
   }

  ngOnInit(): void {
  }

  placeOrder(){
    let checkCartData = JSON.parse(localStorage.getItem('cartData'))
    
    let finalData : any[];
    for(let i=0; i< checkCartData.length; i++){
      let cartData = {
          "itemID": checkCartData[i].id,
          "itemQuantity": checkCartData[i].itemQuantity
      }
      finalData.push(cartData);
    }
    this.api.sendCartData(finalData).subscribe(x => console.log(x))

  }

}
