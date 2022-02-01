import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { APIService } from '../api.service';

@Component({
  selector: 'app-payment',
  templateUrl: './payment.component.html',
  styleUrls: ['./payment.component.scss'],
})
export class PaymentComponent implements OnInit {
  public amount: string;
  public to: string;

  constructor(private route: ActivatedRoute, private api: APIService) {
    this.amount = '';
    this.to = '';
  }

  ngOnInit() {
    this.getQueryParams();
    // this.to = this.route.snapshot.paramMap.get('to');
    // this.amount = this.route.snapshot.paramMap.get('amount');
  }
  sendTransaction() {
    this.api
      .sendTransaction(this.to, this.amount, 'transfer')
      .subscribe((data) => {
        console.log(data);
      });
  }

  getQueryParams() {
    // console.log(this.route.snapshot.paramMap);
    // this.to = this.route.snapshot.paramMap.get('to');

    this.route.queryParams.subscribe((params) => {
      this.amount = params['amount'];
      this.to = params['to'];
    });
  }
}
