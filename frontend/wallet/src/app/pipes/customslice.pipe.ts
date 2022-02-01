import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'customslice',
})
export class CustomSlice implements PipeTransform {
  transform(val: string, length: number): string {
    return val.length > length ? `${val.substring(0, length)} ...` : val;
  }
}
