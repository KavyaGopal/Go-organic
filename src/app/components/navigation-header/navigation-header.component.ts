import { Component, OnInit } from '@angular/core';
import { BreakpointObserver, Breakpoints } from '@angular/cdk/layout';
import { Observable } from 'rxjs';
import { map, shareReplay } from 'rxjs/operators';
import { CartService} from '../../services/cart.service';



@Component({
  selector: 'app-navigation-header',
  templateUrl: './navigation-header.component.html',
  styleUrls: ['./navigation-header.component.css']
})
export class NavigationHeaderComponent implements OnInit{

  myBreakpoint = '(max-width: 900px)';
  cartItem: number = 0;
  isHandset$: Observable<boolean> = this.breakpointObserver.observe(Breakpoints.Handset)
    .pipe(
      map(result => result.matches),
      shareReplay()
    );

    isTab$: Observable<boolean> = this.breakpointObserver.observe([ Breakpoints.XSmall , Breakpoints.Small , Breakpoints.Medium,     Breakpoints.Tablet , Breakpoints.Handset, 
       this.myBreakpoint, Breakpoints.TabletPortrait , Breakpoints.HandsetLandscape ,])
    .pipe(
      map(result => result.matches),
      shareReplay()
    );
    // isTab$: Observable<boolean> = this.breakpointObserver.observe(this.myBreakpoint)
    // .pipe(
    //   map(result => result.matches),
    //   shareReplay()
    // );
    constructor(
      private breakpointObserver: BreakpointObserver, private cartService: CartService) {
        this.cartService.cartSubject.subscribe((data)=> {
          this.cartItem = data;
        })
      }
  
      ngOnInit(): void {
          this.cartItemFunc();
      }
      cartItemFunc(){
        if(localStorage.getItem('cartData') != null){
          var cartCount = JSON.parse(localStorage.getItem('cartData'));
          console.log(cartCount.length);
          this.cartItem = cartCount.length;
        }
      }

}
