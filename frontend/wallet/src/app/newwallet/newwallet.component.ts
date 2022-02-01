import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { APIService } from '../api.service';

@Component({
  selector: 'app-newwallet',
  templateUrl: './newwallet.component.html',
  styleUrls: ['./newwallet.component.scss']
})
export class NewwalletComponent implements OnInit {

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
