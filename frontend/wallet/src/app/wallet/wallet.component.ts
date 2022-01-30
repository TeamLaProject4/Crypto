import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { APIService } from '../api.service';

@Component({
  selector: 'app-wallet',
  templateUrl: './wallet.component.html',
  styleUrls: ['./wallet.component.scss']
})
export class WalletComponent implements OnInit {

  public mnemonic = [];

  constructor(private http: HttpClient,
              private api: APIService) { }

  ngOnInit(): void {
    console.log(this.getMnemonic());
  }

  getMnemonic() {
    this.api.getMnemonic()
    .subscribe((data: never[]) => this.mnemonic = data)
  }
}
