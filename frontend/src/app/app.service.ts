import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable()
export class ConfigService {
  constructor(private http: HttpClient) { }

  addCustomer(data) {
   // console.log('data: ', data)
    return this.http.put("http://localhost:8080/customer", data, { observe: 'response' });
  }

  addTransaction(data) {
    //console.log('data: ', data)
    return this.http.put("http://localhost:8080/transaction", data, { observe: 'response' });
  }

  getCustomers() {
    return this.http.get("http://localhost:8080/customers", { observe: 'response' });
  }

  getCustomer(id) {
    return this.http.get("http://localhost:8080/customer/" + id, { observe: 'response' });
  }

}
