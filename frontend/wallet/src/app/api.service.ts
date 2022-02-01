import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class APIService {
  constructor(private httpClient: HttpClient) {}

  IP = 'http://192.168.178.111';
  PORT = '50929';

  getPublicKey(): Observable<any> {
    return this.httpClient.get(
      this.IP + ':' + this.PORT + '/frontend/publickey'
    );
  }

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

  sendTransaction(
    receiverPk: string,
    amount: string,
    transactionType: string
  ): Observable<any> {
    return this.httpClient.post(
      this.IP + ':' + this.PORT + '/frontend/transaction',
      {
        recieverPublicKey: receiverPk,
        amount: amount,
        transactionType: transactionType,
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
    console.log(accountnumber);
    return this.httpClient.get(
      this.IP +
        ':' +
        this.PORT +
        '/frontend/transactions?publicKey=' +
        accountnumber
    );

    // let queryParams = new HttpParams().append('publicKey', accountnumber);
    // return this.httpClient.get(
    //   this.IP + ':' + this.PORT + '/frontend/transactions',
    //   { params: queryParams }
    // );
  }

  getBlocklength(accountnumber: string): Observable<any> {
    let queryParams = new HttpParams().append('publicKey', accountnumber);
    return this.httpClient.get(
      this.IP + ':' + this.PORT + '/frontend/transactions',
      { params: queryParams }
    );
  }
}
