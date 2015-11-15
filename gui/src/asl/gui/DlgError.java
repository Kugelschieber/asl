/*
 * The MIT License
 *
 * Copyright 2015 Ozan Egitmen.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 */
package asl.gui;

public class DlgError extends javax.swing.JDialog {

    public DlgError(java.awt.Frame parent, boolean modal, String errorMessage) {
        super(parent, modal);
        initComponents();
        lblError.setText(errorMessage);
    }

    private void initComponents() {
        lblError = new javax.swing.JLabel();
        lblTitle = new javax.swing.JLabel();

        setDefaultCloseOperation(javax.swing.WindowConstants.DISPOSE_ON_CLOSE);
        setTitle("ERROR");
        setIconImage(null);
        setMinimumSize(new java.awt.Dimension(380, 150));
        setResizable(false);
        setType(java.awt.Window.Type.POPUP);

        lblError.setFont(new java.awt.Font("Segoe UI Light", 0, 16));
        lblError.setHorizontalAlignment(javax.swing.SwingConstants.CENTER);
        lblError.setText("Some error");

        lblTitle.setFont(new java.awt.Font("Segoe UI Light", 0, 16));
        lblTitle.setHorizontalAlignment(javax.swing.SwingConstants.CENTER);
        lblTitle.setText("asl.exe has encountered an error:");
        lblTitle.setToolTipText("");

        javax.swing.GroupLayout layout = new javax.swing.GroupLayout(getContentPane());
        getContentPane().setLayout(layout);
        layout.setHorizontalGroup(layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING).addComponent(lblTitle, javax.swing.GroupLayout.DEFAULT_SIZE, 380, Short.MAX_VALUE).addComponent(lblError, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, Short.MAX_VALUE));
        layout.setVerticalGroup(layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING).addGroup(layout.createSequentialGroup().addContainerGap().addComponent(lblTitle).addGap(18, 18, 18).addComponent(lblError).addGap(27, 27, 27)));

        pack();
    }

    // Variables declaration - do not modify
    private javax.swing.JLabel lblError;
    private javax.swing.JLabel lblTitle;
    // End of variables declaration
}