import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { Component, NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { RouterModule } from '@angular/router'
import { HomeComponent } from './home.component';
import { BreadcrumbComponent } from './breadcrumb.component';
import { AddCustomerComponent } from './add-customer.component';
import { CustomerDetailComponent } from './customer-detail.component';
import { CustomerAddTransactionComponent } from './customer-add-transaction.component'
@Component({
  selector: 'my-app',
  template: `
  <breadcrumb></breadcrumb>
    <router-outlet></router-outlet>
  `,
})
export class App {
  name: string;
  constructor() {
    this.name = 'Angular2'
  }
}

const routes = [
  {
    path: 'homepage',
    component: HomeComponent,
  },
  {
    path: "addCustomer",
    component: AddCustomerComponent
  },
  {
    path: "customer-detail/:id",
    component: CustomerDetailComponent
  },
  {
    path: "customer-add-transaction/:id",
    component: CustomerAddTransactionComponent
  },
  { path: '', redirectTo: 'homepage', pathMatch: 'full' }
]
@NgModule({
  declarations: [
    App,
    BreadcrumbComponent,
    HomeComponent,
    AddCustomerComponent,
    CustomerDetailComponent,
    CustomerAddTransactionComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    HttpClientModule,
    RouterModule.forRoot(routes),
    FormsModule,
    ReactiveFormsModule
  ],
  providers: [],
  bootstrap: [App]
})
export class AppModule { }
