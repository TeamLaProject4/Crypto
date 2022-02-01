import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class APIService {
  constructor(private httpClient: HttpClient) {}

  IP = 'http://10.51.60.59';
  PORT = '64561';

  getMnemonic(): Observable<any> {
    return this.httpClient.get(
      this.IP + ':' + this.PORT + '/frontend/getMnemonic'
    );
  }

  confirmMnemonic(mnemonic: string): Observable<any> {
    return this.httpClient.post(
      this.IP + ':' + this.PORT + '/frontend/confirmMnemonic',
      {
        mnemonic: mnemonic,
      }
    );
  }

  getBalance(accountnumber: string): Observable<any> {
    let queryParams = new HttpParams().append('publicKey', accountnumber);
    return this.httpClient.get(
      this.IP + ':' + this.PORT + '/frontend/balance',
      { params: queryParams }
    );
  }

  getTransactions(accountnumber: string): Observable<any> {
    let queryParams = new HttpParams().append('publicKey', accountnumber);
    return this.httpClient.get(
      this.IP + ':' + this.PORT + '/frontend/transactions',
      { params: queryParams }
    );
  }

  getBlocklength(accountnumber: string): Observable<any> {
    let queryParams = new HttpParams().append('publicKey', accountnumber);
    return this.httpClient.get(
      this.IP + ':' + this.PORT + '/frontend/transactions',
      { params: queryParams }
    );
  }
}
