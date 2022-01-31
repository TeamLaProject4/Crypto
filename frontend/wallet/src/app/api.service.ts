import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class APIService {


  constructor(private httpClient: HttpClient) { }

  getMnemonic(): Observable<any>{
    return this.httpClient.get("http://localhost:8080/getMnemonic")
  }

  
  confirmMnemonic(mnemonic: string): Observable<any>{
    const httpOptions = {headers: new HttpHeaders({ 
      'Access-Control-Allow-Origin':'*',
    })}

    return this.httpClient.post("http://localhost:8080/confirmMnemonic", //+ `/credentials`, 
    {
      "mnemonic": mnemonic,
    });


}
}
