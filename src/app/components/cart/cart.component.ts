import { Component, OnInit , ViewChild} from '@angular/core';
import {MatPaginator, MatPaginatorModule} from '@angular/material/paginator';
import { MatTableDataSource } from '@angular/material/table';
import { MatSort } from '@angular/material/sort';
import * as groceries from '../../jsonData/groceries.json';
import { ValueConverter } from '@angular/compiler/src/render3/view/template';

@Component({
  selector: 'app-cart',
  templateUrl: './cart.component.html',
  styleUrls: ['./cart.component.css']
})
export class CartComponent implements OnInit {

  displayedColumns: string[] = ['id', 'itemName','imgSrc', 'itemCost', 'itemDesc', 
                    'itemSubtotal','star'];
  dataSource = new MatTableDataSource<cartItem>();
  getCartDetails: any = [];
  total: number = 0;

@ViewChild(MatPaginator) paginator: MatPaginator;
@ViewChild(MatSort) sort: MatSort;
  constructor() {
  //  localStorage.getItem('cartData');;
    
    // this.dataSource =  new MatTableDataSource<cartItem>(ELEMENT_DATA);
   }

  ngOnInit(): void {
    this.cartDetails();
    this.getTotal();
  }

  cartDetails(){
    if(localStorage.getItem('cartData')){
      this.getCartDetails = JSON.parse(localStorage.getItem('cartData'));
      console.log(this.getCartDetails);
      this.dataSource = new MatTableDataSource<cartItem>(this.getCartDetails);
    }
  }

  getTotal(){
    if(localStorage.getItem('cartData')){
      this.getCartDetails = JSON.parse(localStorage.getItem('cartData'));
      this.total =  this.getCartDetails.reduce((acc, val)=>{
        return acc + (val.itemCost * val.itemQuantity) //add quantity in schema
      },0)
      ;
    }
  }

  decQty(item){
    for(let i=0; i<this.getCartDetails.length; i++){
      if(this.getCartDetails[i].id == item.id){
        if(item.itemQuantity !=1) //condition to add only 5 items
        this.getCartDetails[i].itemQuantity = parseInt(item.itemQuantity) - 1; 
        //this.getCartDetails[i] = this.getCartDetails.filter((x)=> { return x.id !== item.id})
        // this.getCartDetails.splice(i,1);
      }

    }
    localStorage.setItem('cartData',JSON.stringify(this.getCartDetails));
    this.getTotal();
    window.location.reload()

  }
  incQty(item){
    for(let i=0; i<this.getCartDetails.length; i++){
      if(this.getCartDetails[i].id == item.id){
        if(item.itemQuantity != item.itemInventory) //condition to add only 5 items
        this.getCartDetails[i].itemQuantity = parseInt(item.itemQuantity) + 1; 
      }

    }
    localStorage.setItem('cartData',JSON.stringify(this.getCartDetails))
    this.getTotal();
    window.location.reload()
  }

  delItm(item){
    for(let i=0; i<this.getCartDetails.length; i++){
      if(this.getCartDetails[i].id == item.id){
        this.getCartDetails.splice(i,1);
        // if(item.itemQuantity !=5) //condition to add only 5 items
        // this.getCartDetails[i].itemQuantity = parseInt(item.itemQuantity) + 1; 
      }

    }
    localStorage.setItem('cartData',JSON.stringify(this.getCartDetails))
    this.getTotal();
    window.location.reload()
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

 const ELEMENT_DATA: cartItem[] = [
  {
      "id": 1,
      "imgSrc": "../../../assets/items/rice.jpeg",
      "itemName": "Rice",
      "itemDesc": "Rice is the seed of the grass species Oryza sativa or less commonly Oryza glaberrima.",
      "itemWt": 1000,
      "itemGross": 1000,
      "itemCost": 18
    }, {
      "id": 2,
      "imgSrc": "../../../assets/items/wheat.jpeg",
      "itemName": "Wheat",
      "itemDesc": "Wheat is a member of the grass family that produces a dry, one-seeded fruit commonly called a kernel.",
      "itemWt": 500,
      "itemGross": 500,
      "itemCost": 12
    }, {
      "id": 3,
      "imgSrc": "../../../assets/items/sugar.jpeg",
      "itemName": "Sugar",
      "itemDesc": "Sugar is the generic name for sweet-tasting, soluble carbohydrates, many of which are used in food.",
      "itemWt": 500,
      "itemGross": 500,
      "itemCost": 18
    }, {
      "id": 3,
      "imgSrc": "../../../assets/items/sugar.jpeg",
      "itemName": "Salt",
      "itemDesc": "SAlt is the generic name for salt-tasting, soluble carbohydrates, many of which are used in food. It is salty in nature.",
      "itemWt": 500,
      "itemGross": 500,
      "itemCost": 18
    },{
      "id": 5,
      "imgSrc": "../../../assets/items/chilli.png",
      "itemName": "Chilli Powder",
      "itemDesc": "Chili powder is the dried, pulverized fruit of one or more varieties of chili pepper, sometimes with the addition of other spices.",
      "itemWt": 100,
      "itemGross": 100,
      "itemCost": 8
    }, {
      "id": 6,
      "imgSrc": "../../../assets/items/chilli.png",
      "itemName": "Garam Masala",
      "itemDesc": "Garam masala is a blend of ground spices originating from South Asia.It is common in Indian, Pakistani, Nepalese and Bangladeshi.",
      "itemWt": 200,
      "itemGross": 200,
      "itemCost": 12
    }
];