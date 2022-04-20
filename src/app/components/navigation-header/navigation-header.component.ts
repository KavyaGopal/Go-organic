import { Component, OnInit } from '@angular/core';
import { BreakpointObserver, Breakpoints } from '@angular/cdk/layout';
import { Observable } from 'rxjs';
import { map, shareReplay } from 'rxjs/operators';
import { CartService} from '../../services/cart.service';
import { NavigationEnd, Router} from '@angular/router'



@Component({
  selector: 'app-navigation-header',
  templateUrl: './navigation-header.component.html',
  styleUrls: ['./navigation-header.component.css']
})
export class NavigationHeaderComponent implements OnInit{

  myBreakpoint = '(max-width: 900px)';
  cartItem: number = 0;
  user: string = null;

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
  sectionScroll: any;
    // isTab$: Observable<boolean> = this.breakpointObserver.observe(this.myBreakpoint)
    // .pipe(
    //   map(result => result.matches),
    //   shareReplay()
    // );
    constructor(
      private breakpointObserver: BreakpointObserver, private cartService: CartService, private router: Router) {
        this.cartService.cartSubject.subscribe((data)=> {
          this.cartItem = data;
          
        })
        debugger;
       
        this.cartService.userName.subscribe((data) => this.user = data);
        console.log('User Name: ' + this.user)
      }
  
      ngOnInit(): void {
          this.cartItemFunc();
          this.router.events.subscribe((evt) => {
            if (!(evt instanceof NavigationEnd)) {
              return;
            }
            this.doScroll();
            this.sectionScroll= null;
          });
      
      }
      cartItemFunc(){
        if(localStorage.getItem('cartData') != null){
          var cartCount = JSON.parse(localStorage.getItem('cartData'));
          console.log(cartCount.length);
          this.cartItem = cartCount.length;
        }
      }

      checkLogin(){
        if(this.user != null){
          localStorage.clear();
          window.location.reload()
        }
        this.router.navigate(['/login']);
      }
      internalRoute(page,dst){
        this.sectionScroll=dst;
        this.router.navigate([page], {fragment: dst});
    }
    doScroll() {

      if (!this.sectionScroll) {
        return;
      }
      try {
        var elements = document.getElementById(this.sectionScroll);
  
        elements.scrollIntoView();
      }
      finally{
        this.sectionScroll = null;
      }
    } 
  
     

}
