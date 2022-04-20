import { Component, OnInit } from '@angular/core';
import { NavigationEnd, Router} from '@angular/router'
@Component({
  selector: 'app-checkout',
  templateUrl: './checkout.component.html',
  styleUrls: ['./checkout.component.css']
})
export class CheckoutComponent implements OnInit {

  constructor(private router: Router) { }

  ngOnInit(): void {

    this.router.events.subscribe((evt) => {
      if (!(evt instanceof NavigationEnd)) {
        return;
      }
      // this.doScroll();
      // this.sectionScroll= null;
    });
  }

  redirectPayment(){
    this.router.navigate(['/paypalpayment']);
    window.location.reload() 
  
  }

}
