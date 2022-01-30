import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class APIService {

  constructor(private httpClient: HttpClient) { }

  getMnemonic(): Observable<any>{
    return this.httpClient.get("http://localhost:8080/mnemonic")
  }
}
