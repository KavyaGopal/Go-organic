import { Component, OnInit, ViewChild } from '@angular/core';
import { MatPaginator } from '@angular/material/paginator';
import { MatTableDataSource } from '@angular/material/table';
import { MatSort } from '@angular/material/sort';
import { NavigationEnd, Router} from '@angular/router';
@Component({
  selector: 'app-checkout',
  templateUrl: './checkout.component.html',
  styleUrls: ['./checkout.component.css']
})
export class CheckoutComponent implements OnInit {

  displayedColumns: string[] = ['id', 'itemName', 'imgSrc', 'itemCost', 'itemDesc',
    'itemSubtotal', 'star'];
  dataSource = new MatTableDataSource<cartItem>();
  getCartDetails: any = [];
  total: number = 0;

  @ViewChild(MatPaginator) paginator: MatPaginator;
  @ViewChild(MatSort) sort: MatSort;
  constructor(private router: Router) {
    //  localStorage.getItem('cartData');;

    // this.dataSource =  new MatTableDataSource<cartItem>(ELEMENT_DATA);
  }

  ngOnInit(): void {
    this.cartDetails();
    this.getTotal();

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

  cartDetails() {
    if (localStorage.getItem('cartData')) {
      this.getCartDetails = JSON.parse(localStorage.getItem('cartData'));
      console.log(this.getCartDetails);
      this.dataSource = new MatTableDataSource<cartItem>(this.getCartDetails);
    }
  }

  getTotal() {
    if (localStorage.getItem('cartData')) {
      this.getCartDetails = JSON.parse(localStorage.getItem('cartData'));
      this.total = this.getCartDetails.reduce((acc, val) => {
        return acc + (val.itemCost * val.itemQuantity) //add quantity in schema
      }, 0)
        ;
    }
  }



}
export interface cartItem {
  id: number;
  imgSrc: string;
  itemCost: number;
  itemDesc: string;
  itemGross: number;
  itemName: string;
  itemWt: number;

}
