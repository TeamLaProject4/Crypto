import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-payment',
  templateUrl: './payment.component.html',
  styleUrls: ['./payment.component.scss']
})
export class PaymentComponent implements OnInit {

  public amount: string | null
  public to: string | null
  
  constructor(private route: ActivatedRoute) {
    this.amount = "";
    this.to = "";
  }

ngOnInit() {
  this.getQueryParams();
  // this.to = this.route.snapshot.paramMap.get('to');
  // this.amount = this.route.snapshot.paramMap.get('amount');

}

getQueryParams(){
  // console.log(this.route.snapshot.paramMap);
  // this.to = this.route.snapshot.paramMap.get('to');

  this.route.queryParams.subscribe(params => {
    this.amount= params['amount'];
    this.to= params['to'];
})
}
}