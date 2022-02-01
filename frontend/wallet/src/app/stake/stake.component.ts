import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { APIService } from '../api.service';

@Component({
  selector: 'app-stake',
  templateUrl: './stake.component.html',
  styleUrls: ['./stake.component.scss'],
})
export class StakeComponent implements OnInit {
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

  sendStake() {
    this.api
      .sendTransaction('STAKE', this.amount, 'STAKE')
      .subscribe((data) => {
        console.log(data);
      });
  }

  getQueryParams() {
    // console.log(this.route.snapshot.paramMap);
    // this.to = this.route.snapshot.paramMap.get('to');
    this.route.queryParams.subscribe((params) => {
      this.amount = params['amount'];
    });
  }
}
