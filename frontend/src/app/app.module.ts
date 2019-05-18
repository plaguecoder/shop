import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { Component, NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { RouterModule } from '@angular/router'
import { SidebarComponent } from './side-bar.component';
import { HomeComponent } from './home.component';
import { BreadcrumbComponent } from './breadcrumb.component';
import { AddCustomerComponent } from './add-customer.component';
import { CustomerDetailComponent } from './customer-detail.component';
import { CustomerAddTransactionComponent } from './customer-add-transaction.component'
import { AreasComponent } from './areas.component'
import { AddAreaComponent } from './add-area.component'
@Component({
  selector: 'my-app',
  template: `
  
  <div class="row">
  <div class="col-md-12">
  <breadcrumb></breadcrumb>
  </div>              
</div>
    <div class="row">
      <div class='col-md-3'>
        <sidebar></sidebar>
      </div>
      <div class='col-md-9'>
        <router-outlet></router-outlet>
      </div>
    </div>
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
    path: '',
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
  {
    path: "areas",
    component: AreasComponent
  },
  {
    path: "add-area",
    component: AddAreaComponent
  }
]
@NgModule({
  declarations: [
    App,
    BreadcrumbComponent,
    SidebarComponent,
    HomeComponent,
    AddCustomerComponent,
    CustomerDetailComponent,
    CustomerAddTransactionComponent,
    AreasComponent,
    AddAreaComponent
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
