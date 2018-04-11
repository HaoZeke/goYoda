#!/usr/bin/env python3
"""
Converts arbitrary divs to TeX environments, complete with env options.
Mainly focused on the flashcards package. 

Copyright (C) 2017  Rohit Goswami <rohit1995@mail.ru>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>. 

"""

import panflute as pf


def prepare(doc):
    pass


def action(elem, doc):
    if isinstance(elem, pf.Div):
        return pf.RawBlock('\begin{'+"dsf"+"}"+"\n"+elem._content,format='latex')


def finalize(doc):
    pass


def main(doc=None):
    return pf.run_filter(action,
                         prepare=prepare,
                         finalize=finalize,
                         doc=doc) 


if __name__ == '__main__':
    main()
