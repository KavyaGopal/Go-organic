import { Component } from '@angular/core';
import { BreakpointObserver, Breakpoints } from '@angular/cdk/layout';
import { Observable } from 'rxjs';
import { map, shareReplay } from 'rxjs/operators';

@Component({
  selector: 'app-navigation-header',
  templateUrl: './navigation-header.component.html',
  styleUrls: ['./navigation-header.component.css']
})
export class NavigationHeaderComponent {

  myBreakpoint = '(max-width: 900px)';
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
    private breakpointObserver: BreakpointObserver) {}

}
