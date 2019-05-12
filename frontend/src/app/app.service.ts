import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable()
export class ConfigService {
  constructor(private http: HttpClient) { }

  addCustomer(data) {
    console.log('data: ', data)
    return this.http.put("http://localhost:8080/merchants", data, { observe: 'response' });
  }
}
