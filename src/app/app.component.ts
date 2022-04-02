import { Component } from '@angular/core';
import { cartItem } from './components/cart/cart.component';
import { GetApiService } from './get-api.service';
import { CartService } from './services/cart.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'Go-organic';
 
  
  constructor(
    private api:GetApiService,private cartService: CartService
  ){
   
  }

  //this calls api on page load
  ngOnInit(){
    this.cartService.userName.next(localStorage.getItem('userData'));
    // this.api.apiCall().subscribe(
    //   (data)=>{
    //     console.warn("get api data ", data);
    //   }
    // )

    // this.api.fruitsApiCall().subscribe(
    //   (data)=>{
    //     console.warn("get api data for Fruits", data);
    //   }
    // )

    // this.api.snacksApiCall().subscribe(
    //   (data)=>{
    //     console.warn("get api data for Snacks ", data);
    //   }
    // )
  }
}
