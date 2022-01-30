import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { APIService } from '../api.service';

@Component({
  selector: 'app-wallet',
  templateUrl: './wallet.component.html',
  styleUrls: ['./wallet.component.scss']
})
export class WalletComponent implements OnInit {

  //public mnemonic = [];
  mnemonic = String();

  constructor(private http: HttpClient,
              private api: APIService) { }

  ngOnInit(): void {
    console.log(this.getMnemonic());
  }

  getMnemonic() {
    this.api.getMnemonic()
    .subscribe(data => {
     this.mnemonic = data
     console.log(this.mnemonic);
    });
  }

  confirmWallet(){
    this.api.confirmMnemonic(this.mnemonic)
    .subscribe(data => {
      console.log(data);
    });
  }
}
