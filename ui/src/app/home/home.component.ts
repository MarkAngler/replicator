import { Component } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { HttpClient } from '@angular/common/http';


@Component({
  selector: 'app-root',
  template: `
    <form [formGroup]="form" (ngSubmit)="onSubmit()">
      <div>
        <label for="server">Server:</label>
        <input id="server" formControlName="server">
      </div>
      <div>
        <label for="port">Port:</label>
        <input id="port" formControlName="port">
      </div>
      <div>
        <label for="username">Username:</label>
        <input id="username" formControlName="username">
      </div>
      <div>
        <label for="password">Password:</label>
        <input id="password" formControlName="password" type="password">
      </div>
      <button type="submit">Submit</button>
    </form>
  `
})
export class AppComponent {
  form: FormGroup;

  constructor(private fb: FormBuilder, private http: HttpClient) {
    this.form = this.fb.group({
      server: [''],
      port: [''],
      username: [''],
      password: ['']
    });
  }

  onSubmit() {
    const formData = this.form.value;

    this.http.post('http://localserver:8080/sourceServers', formData).subscribe(
      (response) => {
        console.log('Server response:', response);
      },
      (error) => {
        console.log('Error:', error);
      }
    );
  }
}

